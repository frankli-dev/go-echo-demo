package pq

import (
	"time"

	"github.com/jinzhu/gorm"
)

// TODO: add comment
type Model struct {
	gorm.Model
	ID uint `gorm:"primary_key;auto_increment"`
}

// TODO: add comment
type RenewableResource struct {
	gorm.Model
	CalendarDate                  time.Time `gorm:"column:calendar_date"`
	TotalRenewableEnergyResources float32   `gorm:"column:total_renewable_energy_resources;type:float"`
	InstalledSolarCapacity        float32   `gorm:"column:installed_rooftop_solar_capacity_minus_losses_mw;type:float"`
	TotalRenewableEnergyPurchased int64     `gorm:"column:total_renewable_energy_purchased_annually_kwh;type:numeric"`
	GreenChoiceSales              int64     `gorm:"column:greenchoice_sales_kwh;type:numeric"`
	RenewableEnergyToFuelCharge   int64     `gorm:"column:renewable_energy_to_fuel_charge_kwh;type:numeric"`
	Wind                          float32   `gorm:"column:wind_mw;type:float"`
	UtilityScaleSolar             float32   `gorm:"column:utility_scale_solar_mw;type:float"`
	Biomass                       float32   `gorm:"column:biomass_mw;type:float"`
}
