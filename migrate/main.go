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
			ID: "20230918_add_component_tables_and_listing_records",
			Migrate: func(tx *gorm.DB) error {
				// Create SocketType and related CPU and Mobo tables
				if err := tx.AutoMigrate(&models.SocketType{}, &models.CPU{}, &models.Mobo{}, &models.GPU{}); err != nil {
					return err
				}

				// Create new component tables (PowerSupply, Memory, Case, Storage)
				if err := tx.AutoMigrate(&models.PowerSupply{}, &models.Memory{}, &models.Case{}, &models.Storage{}); err != nil {
					return err
				}

				// Create new listing record tables for each component
				if err := tx.AutoMigrate(
					&models.CPUListingRecord{},
					&models.MoboListingRecord{},
					&models.GPUListingRecord{},
					&models.PowerSupplyListingRecord{},
					&models.MemoryListingRecord{},
					&models.CaseListingRecord{},
					&models.StorageListingRecord{},
				); err != nil {
					return err
				}

				// Create user and PCBuild tables
				if err := tx.AutoMigrate(&models.User{}, &models.PCBuild{}, &models.Listing{}); err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				// Drop the newly added component and listing record tables on rollback
				if err := tx.Migrator().DropTable(
					"power_supplies",
					"memories",
					"cases",
					"storages",
					"cpu_listing_records",
					"mobo_listing_records",
					"gpu_listing_records",
					"power_supply_listing_records",
					"memory_listing_records",
					"case_listing_records",
					"storage_listing_records",
				); err != nil {
					return err
				}

				// Drop the user and PCBuild tables
				if err := tx.Migrator().DropTable("users", "pc_builds", "listings"); err != nil {
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
