package main

import (
	"bufio"
	"log"
	"os"
	"shared/models"
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Set up the database connection
	dsn := "host=localhost user=powell password=vs3CUREpWord!!1! dbname=pcbb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Open the data file
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("Failed to open data file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentSocketType *models.SocketType
	var currentCPUs []models.CPU
	var currentMobos []models.Mobo
	var currentGPUs []models.GPU
	parsingCPUs := false
	parsingMobos := false
	parsingGPUs := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Parse SocketType
		if strings.HasPrefix(line, "SocketType:") {
			// When starting a new SocketType block, save any pending CPUs or Mobos
			if len(currentCPUs) > 0 {
				saveCPUs(db, currentCPUs)
				currentCPUs = []models.CPU{}
			}
			if len(currentMobos) > 0 {
				saveMobos(db, currentMobos)
				currentMobos = []models.Mobo{}
			}

			// Parse the new SocketType
			socketTypeName := parseStringValue(line)
			var existingSocketType models.SocketType
			if err := db.Where("name = ?", socketTypeName).First(&existingSocketType).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					currentSocketType = &models.SocketType{Name: socketTypeName}
					db.Create(currentSocketType)
				} else {
					log.Fatalf("Failed to check for existing SocketType: %v", err)
				}
			} else {
				currentSocketType = &existingSocketType
				log.Printf("Skipping SocketType %v as it already exists", socketTypeName)
			}
			parsingCPUs, parsingMobos, parsingGPUs = false, false, false
		}

		// Parse RamType
		if strings.HasPrefix(line, "RamType:") {
			ramType := parseIntValue(line)
			currentSocketType.RamType = ramType
			db.Save(currentSocketType)
		}

		// Start parsing CPUs
		if strings.HasPrefix(line, "CPUs:") {
			parsingCPUs, parsingMobos, parsingGPUs = true, false, false
			currentCPUs = []models.CPU{}
		}

		// Start parsing Mobos
		if strings.HasPrefix(line, "Mobos:") {
			parsingCPUs, parsingMobos, parsingGPUs = false, true, false
			currentMobos = []models.Mobo{}
		}

		// Start parsing GPUs
		if strings.HasPrefix(line, "GPUs:") {
			parsingCPUs, parsingMobos, parsingGPUs = false, false, true
			currentGPUs = []models.GPU{}
		}

		// Parse CPU details if within CPUs block
		if parsingCPUs && strings.HasPrefix(line, "- Name:") {
			cpuName := parseStringValue(line)
			cpu := models.CPU{
				SocketTypeID: currentSocketType.ID,
				Name:         cpuName,
			}
			currentCPUs = append(currentCPUs, cpu)
		}

		if parsingCPUs && len(currentCPUs) > 0 && strings.HasPrefix(line, "GenSeries:") {
			currentCPUs[len(currentCPUs)-1].GenSeries = parseStringValue(line)
		}

		if parsingCPUs && len(currentCPUs) > 0 && strings.HasPrefix(line, "Brand:") {
			currentCPUs[len(currentCPUs)-1].Brand = parseStringValue(line)
		}

		if parsingCPUs && len(currentCPUs) > 0 && strings.HasPrefix(line, "Family:") {
			currentCPUs[len(currentCPUs)-1].Family = parseStringValue(line)
		}

		if parsingCPUs && len(currentCPUs) > 0 && strings.HasPrefix(line, "PowerDraw:") {
			currentCPUs[len(currentCPUs)-1].PowerDraw = parseFloatValue(line)
		}

		if parsingCPUs && len(currentCPUs) > 0 && strings.HasPrefix(line, "AvgMarketPrice:") {
			currentCPUs[len(currentCPUs)-1].AvgMarketPrice = parseFloatValue(line)
		}

		if parsingCPUs && len(currentCPUs) > 0 && strings.HasPrefix(line, "PerformanceRating:") {
			currentCPUs[len(currentCPUs)-1].PerformanceRating = parseFloatValue(line)
		}

		if parsingCPUs && len(currentCPUs) > 0 && strings.HasPrefix(line, "SupportLifetime:") {
			currentCPUs[len(currentCPUs)-1].SupportLifetime = parseFloatValue(line)
		}

		// Parse Mobo details if within Mobos block
		if parsingMobos && strings.HasPrefix(line, "- Name:") {
			moboName := parseStringValue(line)
			mobo := models.Mobo{
				SocketTypeID: currentSocketType.ID,
				Name:         moboName,
			}
			currentMobos = append(currentMobos, mobo)
		}

		if parsingMobos && len(currentMobos) > 0 && strings.HasPrefix(line, "Overclocking:") {
			currentMobos[len(currentMobos)-1].Overclocking = parseBoolPointer(line)
		}

		if parsingMobos && len(currentMobos) > 0 && strings.HasPrefix(line, "PCIELaneSupport:") {
			currentMobos[len(currentMobos)-1].PCIELaneSupport = parseIntValue(line)
		}

		if parsingMobos && len(currentMobos) > 0 && strings.HasPrefix(line, "PowerConstrainedCPUs:") {
			currentMobos[len(currentMobos)-1].PowerConstrainedCPUs = parseStringValue(line)
		}

		// Parse GPU details if within GPUs block
		if parsingGPUs && strings.HasPrefix(line, "- Name:") {
			gpuName := parseStringValue(line)
			gpu := models.GPU{
				Name: gpuName,
			}
			currentGPUs = append(currentGPUs, gpu)
		}

		if parsingGPUs && len(currentGPUs) > 0 && strings.HasPrefix(line, "Brand:") {
			currentGPUs[len(currentGPUs)-1].Brand = parseStringValue(line)
		}

		if parsingGPUs && len(currentGPUs) > 0 && strings.HasPrefix(line, "PowerConnection:") {
			currentGPUs[len(currentGPUs)-1].PowerConnection = parseStringValue(line)
		}

		if parsingGPUs && len(currentGPUs) > 0 && strings.HasPrefix(line, "PowerDraw:") {
			currentGPUs[len(currentGPUs)-1].PowerDraw = parseFloatValue(line)
		}

		if parsingGPUs && len(currentGPUs) > 0 && strings.HasPrefix(line, "PerformanceRating:") {
			currentGPUs[len(currentGPUs)-1].PerformanceRating = parseFloatValue(line)
		}

		if parsingGPUs && len(currentGPUs) > 0 && strings.HasPrefix(line, "PCIELaneRequirement:") {
			currentGPUs[len(currentGPUs)-1].PCIELaneRequirement = parseIntValue(line)
		}

		if parsingGPUs && len(currentGPUs) > 0 && strings.HasPrefix(line, "SupportLifetime:") {
			currentGPUs[len(currentGPUs)-1].SupportLifetime = parseFloatValue(line)
		}
	}

	// Save any remaining CPUs, Mobos, or GPUs after the loop
	if len(currentCPUs) > 0 {
		saveCPUs(db, currentCPUs)
	}
	if len(currentMobos) > 0 {
		saveMobos(db, currentMobos)
	}
	if len(currentGPUs) > 0 {
		saveGPUs(db, currentGPUs)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	log.Println("Data loaded successfully.")
}

// Save CPUs to the database
func saveCPUs(db *gorm.DB, cpus []models.CPU) {
	for _, cpu := range cpus {
		var existingCPU models.CPU
		if err := db.Where("name = ?", cpu.Name).First(&existingCPU).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&cpu)
			} else {
				log.Fatalf("Failed to check for existing CPU: %v", err)
			}
		} else {
			log.Printf("Skipping CPU %v as it already exists in the database", cpu.Name)
		}
	}
}

// Save Mobos to the database
func saveMobos(db *gorm.DB, mobos []models.Mobo) {
	for _, mobo := range mobos {
		var existingMobo models.Mobo
		if err := db.Where("name = ?", mobo.Name).First(&existingMobo).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&mobo)
			} else {
				log.Fatalf("Failed to check for existing Mobo: %v", err)
			}
		} else {
			log.Printf("Skipping Mobo %v as it already exists in the database", mobo.Name)
		}
	}
}

// Save GPUs to the database
func saveGPUs(db *gorm.DB, gpus []models.GPU) {
	for _, gpu := range gpus {
		var existingGPU models.GPU
		if err := db.Where("name = ?", gpu.Name).First(&existingGPU).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&gpu)
			} else {
				log.Fatalf("Failed to check for existing GPU: %v", err)
			}
		} else {
			log.Printf("Skipping GPU %v as it already exists in the database", gpu.Name)
		}
	}
}

// Helper functions to handle string, float, and array parsing

func parseStringValue(line string) string {
	return strings.TrimSpace(strings.Split(line, ":")[1])
}

func parseIntValue(line string) int {
	val, _ := strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[1]))
	return val
}

func parseFloatValue(line string) float64 {
	val, _ := strconv.ParseFloat(strings.TrimSpace(strings.Split(line, ":")[1]), 64)
	return val
}

func parseBoolPointer(line string) *bool {
	val := strings.TrimSpace(strings.Split(line, ":")[1])
	booleanValue := val == "true"
	return &booleanValue
}
