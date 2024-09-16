package main

import (
	"log"
	"shared/models" // Import the models from your shared package

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connect to the database
	dsn := "host=localhost user=powell password=vs3CUREpWord!!1! dbname=pcbb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Define the migration with the latest changes to models
	migrations := []*gormigrate.Migration{
		{
			ID: "20230916_create_tables",
			Migrate: func(tx *gorm.DB) error {
				// Create socket_types table first
				if err := tx.AutoMigrate(&models.SocketType{}); err != nil {
					return err
				}
				// Create cpus table and other dependent tables
				if err := tx.AutoMigrate(&models.CPU{}, &models.GPU{}, &models.Mobo{}); err != nil {
					return err
				}
				// Remove CPUPerformanceGroup table (if exists)
				return tx.Migrator().DropTable("cpu_performance_groups")
			},
			Rollback: func(tx *gorm.DB) error {
				// Rollback: drop the new CPU fields and other changes
				return tx.Migrator().DropTable("mobos", "gpus", "cpus", "socket_types")
			},
		},
		{
			// New migration for updating the CPU table to include the new fields
			ID: "20230916_update_cpu_table",
			Migrate: func(tx *gorm.DB) error {
				// Update CPU table to include new fields (GenSeries, Brand, Family)
				return tx.AutoMigrate(&models.CPU{})
			},
			Rollback: func(tx *gorm.DB) error {
				// Rollback: remove the new fields added to the CPU table
				if err := tx.Migrator().DropColumn(&models.CPU{}, "gen_series"); err != nil {
					return err
				}
				if err := tx.Migrator().DropColumn(&models.CPU{}, "brand"); err != nil {
					return err
				}
				return tx.Migrator().DropColumn(&models.CPU{}, "family")
			},
		},
		{
			// New migration for updating the GPU table to include new fields
			ID: "20230916_update_gpu_table",
			Migrate: func(tx *gorm.DB) error {
				// Update GPU table to include new fields (Name, Brand, PowerConnection)
				return tx.AutoMigrate(&models.GPU{})
			},
			Rollback: func(tx *gorm.DB) error {
				// Rollback: remove the new fields added to the GPU table
				if err := tx.Migrator().DropColumn(&models.GPU{}, "name"); err != nil {
					return err
				}
				if err := tx.Migrator().DropColumn(&models.GPU{}, "brand"); err != nil {
					return err
				}
				return tx.Migrator().DropColumn(&models.GPU{}, "power_connection")
			},
		},
		{
			// New migration for making the Overclocking field nullable
			ID: "20230917_make_overclocking_nullable",
			Migrate: func(tx *gorm.DB) error {
				// Modify the Overclocking field to be nullable
				return tx.Migrator().AlterColumn(&models.Mobo{}, "overclocking")
			},
			Rollback: func(tx *gorm.DB) error {
				// Rollback: make Overclocking non-nullable again
				return tx.Migrator().AlterColumn(&models.Mobo{}, "overclocking")
			},
		},
		{
			ID: "20230914_add_name_to_mobo",
			Migrate: func(tx *gorm.DB) error {
				// Ensure the Mobo model has a Name field (single string now)
				return tx.AutoMigrate(&models.Mobo{})
			},
			Rollback: func(tx *gorm.DB) error {
				// Rollback: remove the name field from Mobo table (if necessary)
				return tx.Migrator().DropColumn(&models.Mobo{}, "name")
			},
		},
		{
			ID: "20230915_update_cpu_name_to_string",
			Migrate: func(tx *gorm.DB) error {
				return tx.Migrator().AlterColumn(&models.CPU{}, "name")
			},
			Rollback: func(tx *gorm.DB) error {
				// Rollback logic if necessary
				return tx.Migrator().AlterColumn(&models.CPU{}, "name")
			},
		},
		{
			ID: "20230915_add_gen_series_brand_family_to_cpus",
			Migrate: func(tx *gorm.DB) error {
				// Add new fields to the CPU table
				return tx.AutoMigrate(&models.CPU{})
			},
			Rollback: func(tx *gorm.DB) error {
				// Rollback: remove the new fields
				if err := tx.Migrator().DropColumn(&models.CPU{}, "gen_series"); err != nil {
					return err
				}
				if err := tx.Migrator().DropColumn(&models.CPU{}, "brand"); err != nil {
					return err
				}
				return tx.Migrator().DropColumn(&models.CPU{}, "family")
			},
		},
		{
			ID: "20230915_drop_names_column_from_mobos",
			Migrate: func(tx *gorm.DB) error {
				// Drop the `names` column from the `mobos` table
				if err := tx.Migrator().DropColumn(&models.Mobo{}, "names"); err != nil {
					return err
				}

				// Ensure that the `mobos` table structure is consistent with the updated Go model
				return tx.AutoMigrate(&models.Mobo{})
			},
			Rollback: func(tx *gorm.DB) error {
				// Rollback: add the `names` column back to the `mobos` table
				if err := tx.Migrator().AddColumn(&models.Mobo{}, "names"); err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "20230917_create_sales_tables",
			Migrate: func(tx *gorm.DB) error {
				// AutoMigrate to create the new sales tables for CPU, Mobo, and GPU
				if err := tx.AutoMigrate(&models.CPUSaleRecord{}, &models.MoboSaleRecord{}, &models.GPUSaleRecord{}); err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				// Drop the sales tables in case of rollback
				if err := tx.Migrator().DropTable("cpu_sale_records", "mobo_sale_records", "gpu_sale_records"); err != nil {
					return err
				}
				return nil
			},
		},
	}

	// Initialize gormigrate with version control
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)

	// Run migrations
	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migrations applied successfully.")
}
