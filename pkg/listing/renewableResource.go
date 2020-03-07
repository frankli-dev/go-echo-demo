package listing

// TODO: add comment
type RenewableResource struct {
	CalendarDate                  string
	TotalRenewableEnergyResources float32
	InstalledSolarCapacity        float32
	TotalRenewableEnergyPurchased int64
	GreenChoiceSales              int64
	RenewableEnergyToFuelCharge   int64
	Wind                          float32
	UtilityScaleSolar             float32
	Biomass                       float32
}
