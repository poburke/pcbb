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
	SocketTypeID      uint               `gorm:"not null" json:"socket_type_id"`          // Foreign key to relate to SocketType
	Name              string             `gorm:"type:text" json:"name"`                   // Name of the CPU
	GenSeries         string             `gorm:"type:text" json:"gen_series"`             // CPU generation or series (e.g., 10th Gen, Ryzen 5000)
	Brand             string             `gorm:"type:text" json:"brand"`                  // CPU brand (e.g., Intel, AMD)
	Family            string             `gorm:"type:text" json:"family"`                 // CPU family (e.g., Core, Ryzen)
	PowerDraw         float64            `gorm:"type:float8" json:"power_draw"`           // Power draw in watts
	AvgMarketPrice    float64            `gorm:"type:float8" json:"avg_market_price"`     // Average market price of the CPU
	PerformanceRating float64            `gorm:"type:float8" json:"performance_rating"`   // Performance rating of the CPU
	SupportLifetime   float64            `gorm:"type:float8" json:"support_lifetime"`     // Expected support lifetime in years
	ListingRecords    []CPUListingRecord `gorm:"foreignKey:CPUId" json:"listing_records"` // Listing records for this CPU
}

// Mobo represents the motherboard data structure
type Mobo struct {
	gorm.Model
	SocketTypeID         uint                `gorm:"not null" json:"socket_type_id"`           // Foreign key to relate to SocketType
	Name                 string              `gorm:"type:text" json:"name"`                    // Name of the Mobo
	GenSeries            []string            `gorm:"type:text[]" json:"gen_series"`            // Array of compatible CPU generations or series
	Overclocking         *bool               `gorm:"type:bool" json:"overclocking"`            // Whether the motherboard supports overclocking (nullable)
	PCIELaneSupport      int                 `gorm:"type:int" json:"pcie_lane_support"`        // Number of PCIe lanes supported
	PowerConstrainedCPUs string              `gorm:"type:text" json:"power_constrained_cpus"`  // List of power-constrained CPUs (nullable)
	ListingRecords       []MoboListingRecord `gorm:"foreignKey:MoboId" json:"listing_records"` // Listing records for this Mobo
}

// GPU represents the GPU data structure
type GPU struct {
	gorm.Model
	Name                string             `gorm:"type:text" json:"name"`                   // Name of the GPU
	Brand               string             `gorm:"type:text" json:"brand"`                  // GPU brand (e.g., NVIDIA, AMD)
	PowerConnection     string             `gorm:"type:text" json:"power_connection"`       // Type of power connection required (e.g., 6-pin, 8-pin)
	PowerDraw           float64            `gorm:"type:float8" json:"power_draw"`           // Power draw in watts
	PerformanceRating   float64            `gorm:"type:float8" json:"performance_rating"`   // Performance rating of the GPU
	PCIELaneRequirement int                `gorm:"type:int" json:"pcie_lane_requirement"`   // PCIe lane requirement
	SupportLifetime     float64            `gorm:"type:float8" json:"support_lifetime"`     // Expected support lifetime in years
	ListingRecords      []GPUListingRecord `gorm:"foreignKey:GPUId" json:"listing_records"` // Listing records for this GPU
}

// PowerSupply represents the power supply data structure
type PowerSupply struct {
	gorm.Model
	Name             string                     `gorm:"type:text" json:"name"`                           // Name of the PowerSupply
	PowerOutput      float64                    `gorm:"type:float8" json:"power_output"`                 // Power output in watts
	EfficiencyRating string                     `gorm:"type:text" json:"efficiency_rating"`              // Efficiency rating (e.g., 80+ Gold)
	OverallRating    *float64                   `gorm:"type:float8;nullable" json:"overall_rating"`      // Overall rating (nullable)
	ListingRecords   []PowerSupplyListingRecord `gorm:"foreignKey:PowerSupplyId" json:"listing_records"` // Listing records for the PowerSupply
}

// Memory represents the memory data structure
type Memory struct {
	gorm.Model
	Name           string                `gorm:"type:text" json:"name"`                      // Name of the Memory module
	TransferSpeed  float64               `gorm:"type:float8" json:"transfer_speed"`          // Transfer speed in MHz
	Type           string                `gorm:"type:text" json:"type"`                      // Memory type (e.g., DDR4)
	Size           int                   `gorm:"type:int" json:"size"`                       // Size in GB
	Brand          *string               `gorm:"type:text;nullable" json:"brand"`            // Memory brand (nullable)
	ListingRecords []MemoryListingRecord `gorm:"foreignKey:MemoryId" json:"listing_records"` // Listing records for Memory
}

// Case represents the case data structure
type Case struct {
	gorm.Model
	Name           string              `gorm:"type:text" json:"name"`                      // Name of the Case
	FormFactor     string              `gorm:"type:text" json:"form_factor"`               // Form factor (e.g., ATX, Micro-ATX)
	OverallRating  *float64            `gorm:"type:float8;nullable" json:"overall_rating"` // Overall rating (nullable)
	ListingRecords []CaseListingRecord `gorm:"foreignKey:CaseId" json:"listing_records"`   // Listing records for Case
}

// Storage represents the storage data structure
type Storage struct {
	gorm.Model
	Name           string                 `gorm:"type:text" json:"name"`                       // Name of the Storage
	Type           string                 `gorm:"type:text" json:"type"`                       // Storage type (e.g., SSD, HDD)
	Size           int                    `gorm:"type:int" json:"size"`                        // Size in GB
	Bandwidth      *float64               `gorm:"type:float8;nullable" json:"bandwidth"`       // Bandwidth (nullable)
	Generation     *string                `gorm:"type:text;nullable" json:"generation"`        // Generation (e.g., PCIe Gen 4, SATA III) (nullable)
	OverallRating  *float64               `gorm:"type:float8;nullable" json:"overall_rating"`  // Overall rating (nullable)
	ListingRecords []StorageListingRecord `gorm:"foreignKey:StorageId" json:"listing_records"` // Listing records for Storage
}

// User represents a user in the system
type User struct {
	gorm.Model
	Username      string    `gorm:"type:text;unique" json:"username"`        // User's email (from Keycloak)
	LastAccessed  time.Time `gorm:"type:timestamp" json:"last_accessed"`     // Last time the user accessed the site
	SavedBuilds   []PCBuild `gorm:"foreignKey:UserID" json:"saved_builds"`   // User's saved builds
	SavedListings []Listing `gorm:"foreignKey:UserID" json:"saved_listings"` // User's saved eBay listings
}

// PCBuild represents a custom PC build that a user has saved
type PCBuild struct {
	gorm.Model
	UserID        uint   `gorm:"not null" json:"user_id"`                // Foreign key to relate to User
	MoboID        *uint  `gorm:"foreignKey:MoboID" json:"mobo_id"`       // Motherboard selection
	CPUId         *uint  `gorm:"foreignKey:CPUId" json:"cpu_id"`         // CPU selection
	GPUId         *uint  `gorm:"foreignKey:GPUId" json:"gpu_id"`         // GPU selection
	StorageID     *uint  `gorm:"foreignKey:StorageID" json:"storage_id"` // Storage selection
	MemoryID      *uint  `gorm:"foreignKey:MemoryID" json:"memory_id"`   // Memory selection
	PowerSupplyID *uint  `gorm:"foreignKey:PowerSupplyID" json:"psu_id"` // Power supply selection
	CaseID        *uint  `gorm:"foreignKey:CaseID" json:"case_id"`       // Case selection
	BuildName     string `gorm:"type:text" json:"build_name"`            // Name of the build
	BuildNote     string `gorm:"type:text" json:"build_note"`            // User's notes on the build
}

// Listing represents an actual listing (eBay or other source) that users have saved
type Listing struct {
	gorm.Model
	UserID        uint    `gorm:"not null" json:"user_id"`         // Foreign key to relate to User
	ComponentType string  `gorm:"type:text" json:"component_type"` // Type of component (e.g., CPU, GPU)
	ListingURL    string  `gorm:"type:text" json:"listing_url"`    // URL of the listing
	Price         float64 `gorm:"type:float8" json:"price"`        // Listing price
	Seller        string  `gorm:"type:text" json:"seller"`         // Seller information
	Rating        float64 `gorm:"type:float8" json:"rating"`       // Rating for the seller
}

// ListingRecord for CPU
type CPUListingRecord struct {
	gorm.Model
	CPUId        uint      `gorm:"not null" json:"cpu_id"`           // Foreign key to relate to CPU
	ListingPrice float64   `gorm:"type:float8" json:"listing_price"` // Listing price for the CPU
	Volume       int       `gorm:"type:int" json:"volume"`           // Volume of listings
	Day          time.Time `gorm:"type:date" json:"day"`             // Date of the listing
}

// ListingRecord for Mobo
type MoboListingRecord struct {
	gorm.Model
	MoboId       uint      `gorm:"not null" json:"mobo_id"`          // Foreign key to relate to Mobo
	ListingPrice float64   `gorm:"type:float8" json:"listing_price"` // Listing price for the Mobo
	Volume       int       `gorm:"type:int" json:"volume"`           // Volume of listings
	Day          time.Time `gorm:"type:date" json:"day"`             // Date of the listing
}

// ListingRecord for GPU
type GPUListingRecord struct {
	gorm.Model
	GPUId        uint      `gorm:"not null" json:"gpu_id"`           // Foreign key to relate to GPU
	ListingPrice float64   `gorm:"type:float8" json:"listing_price"` // Listing price for the GPU
	Volume       int       `gorm:"type:int" json:"volume"`           // Volume of listings
	Day          time.Time `gorm:"type:date" json:"day"`             // Date of the listing
}

// ListingRecord for PowerSupply
type PowerSupplyListingRecord struct {
	gorm.Model
	PowerSupplyId uint      `gorm:"not null" json:"power_supply_id"`  // Foreign key to relate to PowerSupply
	ListingPrice  float64   `gorm:"type:float8" json:"listing_price"` // Listing price for the PowerSupply
	Volume        int       `gorm:"type:int" json:"volume"`           // Volume of listings
	Day           time.Time `gorm:"type:date" json:"day"`             // Date of the listing
}

// ListingRecord for Memory
type MemoryListingRecord struct {
	gorm.Model
	MemoryId     uint      `gorm:"not null" json:"memory_id"`        // Foreign key to relate to Memory
	ListingPrice float64   `gorm:"type:float8" json:"listing_price"` // Listing price for the Memory
	Volume       int       `gorm:"type:int" json:"volume"`           // Volume of listings
	Day          time.Time `gorm:"type:date" json:"day"`             // Date of the listing
}

// ListingRecord for Case
type CaseListingRecord struct {
	gorm.Model
	CaseId       uint      `gorm:"not null" json:"case_id"`          // Foreign key to relate to Case
	ListingPrice float64   `gorm:"type:float8" json:"listing_price"` // Listing price for the Case
	Volume       int       `gorm:"type:int" json:"volume"`           // Volume of listings
	Day          time.Time `gorm:"type:date" json:"day"`             // Date of the listing
}

// ListingRecord for Storage
type StorageListingRecord struct {
	gorm.Model
	StorageId    uint      `gorm:"not null" json:"storage_id"`       // Foreign key to relate to Storage
	ListingPrice float64   `gorm:"type:float8" json:"listing_price"` // Listing price for the Storage
	Volume       int       `gorm:"type:int" json:"volume"`           // Volume of listings
	Day          time.Time `gorm:"type:date" json:"day"`             // Date of the listing
}
