package models

type Order struct {
	ID                    int                           `json:"id"`
	DryCleaning           *[]DryCleaningOrder           `json:"dry_cleaning,omitempty"`
	HandWash              *HandWashOrder                `json:"hand_wash,omitempty"`
	GeneralLaundryService *[]GeneralLaundryServiceOrder `json:"general_laundry_service,omitempty"`
	IroningServices       *[]IroningServicesOrder       `json:"ironing_services,omitempty"`
	ClothingRepair        *ClothingRepairOrder          `json:"clothing_repair,omitempty"`
	StainRemoval          *[]StainRemovalOrder          `json:"stain_removal,omitempty"`
	TotalAmount           float32                       `json:"total_amount"`
}

type DryCleaningOrder struct {
	ID          int     `json:"id"`
	OrderID     int     `json:"order_id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
	TotalAmount float32 `json:"total_amount"`
}

type HandWashOrder struct {
	ID          int     `json:"id"`
	OrderID     int     `json:"order_id"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
	TotalAmount float32 `json:"total_amount"`
}

type GeneralLaundryServiceOrder struct {
	ID              int     `json:"id"`
	OrderID         int     `json:"order_id"`
	ProductTypeID   int     `json:"product_type_id"`
	ProductTypeName string  `json:"product_type_name"`
	Weight          float32 `json:"weight"`
	Price           float32 `json:"price"`
	TotalAmount     float32 `json:"total_amount"`
}

type IroningServicesOrder struct {
	ID          int     `json:"id"`
	OrderID     int     `json:"order_id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
	TotalAmount float32 `json:"total_amount"`
}

type ClothingRepairOrder struct {
	ID          int     `json:"id"`
	OrderID     int     `json:"order_id"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
	TotalAmount float32 `json:"total_amount"`
}

type StainRemovalOrder struct {
	ID          int     `json:"id"`
	OrderID     int     `json:"order_id"`
	TypeID      int     `json:"type_id"`
	TypeName    string  `json:"type_name"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
	TotalAmount float32 `json:"total_amount"`
}
