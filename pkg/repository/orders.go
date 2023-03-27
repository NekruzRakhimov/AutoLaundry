package repository

import (
	"fmt"
	"github.com/NekruzRakhimov/AutoLaundery/db"
	"github.com/NekruzRakhimov/AutoLaundery/logger"
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/utils"
	"gorm.io/gorm"
)

func SaveOrder(trx *gorm.DB, order *models.Order) error {
	if err := trx.Table("orders").Select("total_amount").Create(order).Error; err != nil {
		trx.Rollback()
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return err
	}

	return nil
}

func UpdateOrder(trx *gorm.DB, order *models.Order) error {
	if err := trx.Table("orders").Select("total_amount", "discount_ids", "mark_up_ids").
		Save(order).Error; err != nil {
		trx.Rollback()
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return err
	}

	return nil
}

func GetAllOrders(date string) (orders []models.Order, err error) {
	sqlQuery := `select id, total_amount
					from orders`
	if date != "" {
		sqlQuery += fmt.Sprintf(" where DATE(created_at) = '%s'", date)
	}

	sqlQuery += " ORDER BY id"

	if err = db.GetDBConn().Raw(sqlQuery).Scan(&orders).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return nil, err
	}

	return
}

func GetOrderByID(id int) (order models.OrderDB, err error) {
	sqlQuery := `select id, total_amount, discount_ids, mark_up_ids
			from orders
			WHERE id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, id).Scan(&order).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return models.OrderDB{}, err
	}

	return
}

func SaveDryCleaningOrder(trx *gorm.DB, orderID int, orders *[]models.DryCleaningOrder) error {
	for i := range *orders {
		(*orders)[i].OrderID = orderID
	}

	if err := trx.Table("dry_cleaning_orders").Omit("product_name").Create(&orders).Error; err != nil {
		trx.Rollback()
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return err
	}

	return nil
}

func SaveHandWashOrder(trx *gorm.DB, orderID int, order *models.HandWashOrder) error {
	order.OrderID = orderID
	if err := trx.Table("hand_wash_orders").Create(&order).Error; err != nil {
		trx.Rollback()
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return err
	}

	return nil
}

func SaveGeneralLaundryServiceOrder(trx *gorm.DB, orderID int, orders *[]models.GeneralLaundryServiceOrder) error {
	for i := range *orders {
		(*orders)[i].OrderID = orderID
	}

	if err := trx.Table("general_laundry_service_orders").Omit("product_type_name").
		Create(&orders).Error; err != nil {
		trx.Rollback()
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return err
	}

	return nil
}

func SaveIroningServicesOrder(trx *gorm.DB, orderID int, orders *[]models.IroningServicesOrder) error {
	for i := range *orders {
		(*orders)[i].OrderID = orderID
	}

	if err := trx.Table("ironing_services_orders").Omit("product_name").
		Create(&orders).Error; err != nil {
		trx.Rollback()
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return err
	}

	return nil
}

func SaveClothingRepairOrder(trx *gorm.DB, orderID int, order *models.ClothingRepairOrder) error {
	order.OrderID = orderID
	if err := trx.Table("clothing_repair_orders").Create(&order).Error; err != nil {
		trx.Rollback()
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return err
	}

	return nil
}

func SaveStainRemovalOrder(trx *gorm.DB, orderID int, orders *[]models.StainRemovalOrder) error {
	for i := range *orders {
		(*orders)[i].OrderID = orderID
	}

	if err := trx.Table("stain_removal_orders").Omit("type_name").
		Create(&orders).Error; err != nil {
		trx.Rollback()
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return err
	}

	return nil
}

func GetDryCleaningOrderByID(orderID int) (order *[]models.DryCleaningOrderDB, err error) {
	sqlQuery := `select dco.id,
					   dco.order_id,
					   dco.product_id,
					   p.name as product_name,
					   dco.quantity,
					   dco.price,
					   dco.total_amount,
					   dco.discount_ids,
					   dco.mark_up_ids
				from dry_cleaning_orders dco
						 join products p on dco.product_id = p.id
				where dco.order_id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, orderID).Scan(&order).Error; err != nil {
		return nil, err
	}

	return
}

func GetHandWashOrderByID(orderID int) (order *models.HandWashOrderDB, err error) {
	sqlQuery := `select hwo.id,
					   hwo.order_id,
					   hwo.quantity,
					   hwo.price,
					   hwo.total_amount,
					   hwo.discount_ids,
					   hwo.mark_up_ids
				from hand_wash_orders hwo
				where order_id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, orderID).Scan(&order).Error; err != nil {
		return nil, err
	}

	return
}

func GetGeneralLaundryServiceOrderByID(orderID int) (order *[]models.GeneralLaundryServiceOrderDB, err error) {
	sqlQuery := `select glso.id,
			   glso.order_id,
			   glso.product_type_id,
				glp.name as product_type_name,
				glso.weight,
				glso.price,
				glso.total_amount,
				glso.discount_ids,
				glso.mark_up_ids
		from general_laundry_service_orders glso
				 join general_laundry_pricing glp on glp.id = glso.product_type_id
		where glso.order_id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, orderID).Scan(&order).Error; err != nil {
		return nil, err
	}

	return
}

func GetIroningServicesOrderByID(orderID int) (order *[]models.IroningServicesOrderDB, err error) {
	sqlQuery := `select iso.id,
					   iso.order_id,
					   iso.product_id,
					   p.name as product_name,
					   iso.quantity,
					   iso.price,
					   iso.total_amount,
					   iso.discount_ids,
					   iso.mark_up_ids
					from ironing_services_orders iso join products p on iso.product_id = p.id
				where order_id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, orderID).Scan(&order).Error; err != nil {
		return nil, err
	}

	return
}

func GetClothingRepairOrderByID(orderID int) (order *models.ClothingRepairOrderDB, err error) {
	sqlQuery := `select cro.id,
					   cro.order_id,
					   cro.quantity,
					   cro.price,
					   cro.total_amount,
					   cro.discount_ids,
					   cro.mark_up_ids
					   from clothing_repair_orders cro
				where order_id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, orderID).Scan(&order).Error; err != nil {
		return nil, err
	}

	return
}

func GetStainRemovalOrderByID(orderID int) (order *[]models.StainRemovalOrderDB, err error) {
	sqlQuery := `select sro.id,
					   sro.order_id,
					   sro.type_id,
					   srp.name as type_name,
					   sro.quantity,
					   sro.price,
					   sro.total_amount,
					   sro.discount_ids,
					   sro.mark_up_ids
					from stain_removal_orders sro
				join stain_removal_pricing srp on sro.type_id = srp.id
				where order_id = ?`
	if err = db.GetDBConn().Raw(sqlQuery, orderID).Scan(&order).Error; err != nil {
		return nil, err
	}

	return
}
