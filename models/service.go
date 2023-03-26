package models

const (
	DryCleaningCode           = "DRY_CLEANING"
	HandWashCode              = "HAND_WASH"
	GeneralLaundryServiceCode = "GENERAL_LAUNDRY_SERVICES"
	IroningServicesCode       = "IRONING_SERVICES"
	ClothingRepairCode        = "CLOTHING_REPAIR"
	StainRemovalCode          = "STAIN_REMOVAL"
)

type Service struct {
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	Code                string  `json:"code"`
	Description         string  `json:"description"`
	HasUniversalPrice   bool    `json:"has_universal_price"`
	Price               float32 `json:"-"`
	UniversalPriceTitle string  `json:"universal_price_title"`
}

type ServicePricing struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
