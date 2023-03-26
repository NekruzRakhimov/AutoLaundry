package models

type Order struct {
	ID                    int                           `json:"id"`
	DryCleaning           *[]DryCleaningOrder           `json:"dry_cleaning"`
	HandWash              *HandWashOrder                `json:"hand_wash"`
	GeneralLaundryService *[]GeneralLaundryServiceOrder `json:"general_laundry_service"`
	IroningServices       *[]IroningServicesOrder       `json:"ironing_services"`
	ClothingRepair        *ClothingRepairOrder          `json:"clothing_repair"`
	StainRemoval          *[]StainRemovalOrder          `json:"stain_removal"`
	TotalAmount           float32                       `json:"total_amount"`
}

type DryCleaningOrder struct {
	ID          int     `json:"id"`
	OrderID     int     `json:"order_id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
}

type HandWashOrder struct {
	ID       int `json:"id"`
	OrderID  int `json:"order_id"`
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

type GeneralLaundryServiceOrder struct {
	ID              int     `json:"id"`
	OrderID         int     `json:"order_id"`
	ProductTypeID   int     `json:"product_type_id"`
	ProductTypeName string  `json:"product_type_name"`
	Weight          float32 `json:"weight"`
	Price           float32 `json:"price"`
}

type IroningServicesOrder struct {
	ID          int     `json:"id"`
	OrderID     int     `json:"order_id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
}

type ClothingRepairOrder struct {
	ID       int     `json:"id"`
	OrderID  int     `json:"order_id"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}

type StainRemovalOrder struct {
	ID       int     `json:"id"`
	OrderID  int     `json:"order_id"`
	TypeID   int     `json:"type_id"`
	TypeName string  `json:"type_name"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}
