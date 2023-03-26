package repository

import (
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
	if err := trx.Table("orders").Select("total_amount").Save(order).Error; err != nil {
		trx.Rollback()
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return err
	}

	return nil
}

func GetAllOrders() (orders []models.Order, err error) {
	if err = db.GetDBConn().Table("orders").Scan(&orders).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return nil, err
	}

	return
}

func GetOrderByID(id int) (order models.Order, err error) {
	if err = db.GetDBConn().Table("orders").Where("id = ?", id).Scan(&order).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return models.Order{}, err
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
