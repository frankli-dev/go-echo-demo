package pq

import (
	"fmt"
	"os"

	"github.com/ivan-marquez/golang-demo/pkg/listing"
	"github.com/jinzhu/gorm"
)

// Storage struct with gorm (postgres) implementation
type Storage struct {
	DB *gorm.DB
}

// GetAllRenewableResources retrieves all renewable resources
func (s *Storage) GetAllRenewableResources() ([]*listing.RenewableResource, error) {
	var resources []*RenewableResource
	var list []*listing.RenewableResource

	if err := s.DB.Find(&resources).Error; err != nil {
		return nil, err
	}

	for _, r := range resources {
		list = append(list, &listing.RenewableResource{
			CalendarDate:                  r.CalendarDate.Format("01/02/2006"),
			TotalRenewableEnergyResources: r.TotalRenewableEnergyResources,
			InstalledSolarCapacity:        r.InstalledSolarCapacity,
			TotalRenewableEnergyPurchased: r.TotalRenewableEnergyPurchased,
			GreenChoiceSales:              r.GreenChoiceSales,
			RenewableEnergyToFuelCharge:   r.RenewableEnergyToFuelCharge,
			Wind:                          r.Wind,
			UtilityScaleSolar:             r.UtilityScaleSolar,
			Biomass:                       r.Biomass,
		})
	}

	return list, nil
}

// NewStorage returns a Storage with postgres connection setup
func NewStorage() (*Storage, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("POSTGRES_USER")
	pw := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbName,
		pw,
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting to database: %v", err)
	}

	s := new(Storage)
	s.DB = db

	return s, nil
}
