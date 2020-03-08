package pq

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Model injects ID, CreatedAt, UpdatedAt, DeletedAt fields to all models
type Model struct {
	gorm.Model
	ID uint `gorm:"primary_key;auto_increment"`
}

// RenewableResource represents a renewable resource record
type RenewableResource struct {
	gorm.Model
	CalendarDate                  time.Time `gorm:"column:calendar_date"`                                               // Fiscal Year
	TotalRenewableEnergyResources float32   `gorm:"column:total_renewable_energy_resources;type:float"`                 // Total Renewable Energy Resources
	InstalledSolarCapacity        float32   `gorm:"column:installed_rooftop_solar_capacity_minus_losses_mw;type:float"` // Installed Rooftop Solar Capacity Minus Losses (MW)
	TotalRenewableEnergyPurchased int64     `gorm:"column:total_renewable_energy_purchased_annually_kwh;type:numeric"`  // Total Renewable Energy Purchased Annually (kWh)
	GreenChoiceSales              int64     `gorm:"column:greenchoice_sales_kwh;type:numeric"`                          // GreenChoice®  Sales (kWh)
	RenewableEnergyToFuelCharge   int64     `gorm:"column:renewable_energy_to_fuel_charge_kwh;type:numeric"`            // Renewable Energy to Fuel Charge (kWh)
	Wind                          float32   `gorm:"column:wind_mw;type:float"`                                          // Wind (MW)
	UtilityScaleSolar             float32   `gorm:"column:utility_scale_solar_mw;type:float"`                           // Utility-Scale Solar (MW)
	Biomass                       float32   `gorm:"column:biomass_mw;type:float"`                                       // Biomass (MW)
}
