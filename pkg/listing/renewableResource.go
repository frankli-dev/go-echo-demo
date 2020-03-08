package listing

// RenewableResource struct represents
// a renewable resource
type RenewableResource struct {
	CalendarDate                  string  // Fiscal Year
	TotalRenewableEnergyResources float32 // Total Renewable Energy Resources
	InstalledSolarCapacity        float32 // Installed Rooftop Solar Capacity Minus Losses (MW)
	TotalRenewableEnergyPurchased int64   // Total Renewable Energy Purchased Annually (kWh)
	GreenChoiceSales              int64   // GreenChoice®  Sales (kWh)
	RenewableEnergyToFuelCharge   int64   // Renewable Energy to Fuel Charge (kWh)
	Wind                          float32 // Wind (MW)
	UtilityScaleSolar             float32 // Utility-Scale Solar (MW)
	Biomass                       float32 // Biomass (MW)
}
