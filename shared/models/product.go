package models

import (
	"time"

	"gorm.io/gorm"
)

// SocketType represents the socket type, related to both CPU and Mobo
type SocketType struct {
	gorm.Model
	Name    string `gorm:"type:text;unique" json:"name"`         // Name of the socket type (e.g., LGA1151, AM4)
	RamType int    `gorm:"type:int" json:"ram_type"`             // RAM type supported by the socket (e.g., DDR4, DDR5)
	CPUs    []CPU  `gorm:"foreignKey:SocketTypeID" json:"cpus"`  // List of CPUs that use this socket type
	Mobos   []Mobo `gorm:"foreignKey:SocketTypeID" json:"mobos"` // List of motherboards that use this socket type
}

// CPU represents the CPU data structure
type CPU struct {
	gorm.Model
	SocketTypeID      uint            `gorm:"not null" json:"socket_type_id"`        // Foreign key to relate to SocketType
	Name              string          `gorm:"type:text" json:"name"`                 // Name of the CPU
	GenSeries         string          `gorm:"type:text" json:"gen_series"`           // CPU generation or series (e.g., 10th Gen, Ryzen 5000)
	Brand             string          `gorm:"type:text" json:"brand"`                // CPU brand (e.g., Intel, AMD)
	Family            string          `gorm:"type:text" json:"family"`               // CPU family (e.g., Core, Ryzen)
	PowerDraw         float64         `gorm:"type:float8" json:"power_draw"`         // Power draw in watts
	AvgMarketPrice    float64         `gorm:"type:float8" json:"avg_market_price"`   // Average market price of the CPU
	PerformanceRating float64         `gorm:"type:float8" json:"performance_rating"` // Performance rating of the CPU
	SupportLifetime   float64         `gorm:"type:float8" json:"support_lifetime"`   // Expected support lifetime in years
	SaleRecords       []CPUSaleRecord `gorm:"foreignKey:CPUId"`                      // Sale records for this CPU
}

// CPUSaleRecord represents sales data for CPUs
type CPUSaleRecord struct {
	gorm.Model
	CPUId     uint      `gorm:"not null" json:"cpu_id"`        // Foreign key to relate to CPU
	SalePrice float64   `gorm:"type:float8" json:"sale_price"` // Sale price for the CPU
	Volume    int       `gorm:"type:int" json:"volume"`        // Volume of sales
	Day       time.Time `gorm:"type:date" json:"day"`          // Date of the sale
}

// Mobo represents the motherboard data structure
type Mobo struct {
	gorm.Model
	SocketTypeID         uint             `gorm:"not null" json:"socket_type_id"`          // Foreign key to relate to SocketType
	Name                 string           `gorm:"type:text" json:"name"`                   // Name of the Mobo
	GenSeries            []string         `gorm:"type:text[]" json:"gen_series"`           // Array of compatible CPU generations or series
	Overclocking         *bool            `gorm:"type:bool" json:"overclocking"`           // Whether the motherboard supports overclocking (nullable)
	PCIELaneSupport      int              `gorm:"type:int" json:"pcie_lane_support"`       // Number of PCIe lanes supported
	PowerConstrainedCPUs string           `gorm:"type:text" json:"power_constrained_cpus"` // List of power-constrained CPUs (nullable, now a single string)
	SaleRecords          []MoboSaleRecord `gorm:"foreignKey:MoboId"`                       // Sale records for this Mobo
}

// MoboSaleRecord represents sales data for motherboards
type MoboSaleRecord struct {
	gorm.Model
	MoboId    uint      `gorm:"not null" json:"mobo_id"`       // Foreign key to relate to Mobo
	SalePrice float64   `gorm:"type:float8" json:"sale_price"` // Sale price for the Mobo
	Volume    int       `gorm:"type:int" json:"volume"`        // Volume of sales
	Day       time.Time `gorm:"type:date" json:"day"`          // Date of the sale
}

// GPU represents the GPU data structure
type GPU struct {
	gorm.Model
	Name                string          `gorm:"type:text" json:"name"`                 // Name of the GPU
	Brand               string          `gorm:"type:text" json:"brand"`                // GPU brand (e.g., NVIDIA, AMD)
	PowerConnection     string          `gorm:"type:text" json:"power_connection"`     // Type of power connection required (e.g., 6-pin, 8-pin)
	PowerDraw           float64         `gorm:"type:float8" json:"power_draw"`         // Power draw in watts
	PerformanceRating   float64         `gorm:"type:float8" json:"performance_rating"` // Performance rating of the GPU
	PCIELaneRequirement int             `gorm:"type:int" json:"pcie_lane_requirement"` // PCIe lane requirement
	SupportLifetime     float64         `gorm:"type:float8" json:"support_lifetime"`   // Expected support lifetime in years
	SaleRecords         []GPUSaleRecord `gorm:"foreignKey:GPUId"`                      // Sale records for this GPU
}

// GPUSaleRecord represents sales data for GPUs
type GPUSaleRecord struct {
	gorm.Model
	GPUId     uint      `gorm:"not null" json:"gpu_id"`        // Foreign key to relate to GPU
	SalePrice float64   `gorm:"type:float8" json:"sale_price"` // Sale price for the GPU
	Volume    int       `gorm:"type:int" json:"volume"`        // Volume of sales
	Day       time.Time `gorm:"type:date" json:"day"`          // Date of the sale
}
