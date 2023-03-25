package models

const (
	DryCleaning           = "DRY_CLEANING"
	HandWash              = "HAND_WASH"
	GeneralLaundryService = "GENERAL_LAUNDRY_SERVICES"
	IroningServices       = "IRONING_SERVICES"
	ClothingRepair        = "CLOTHING_REPAIR"
	StainRemoval          = "STAIN_REMOVAL"
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
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
