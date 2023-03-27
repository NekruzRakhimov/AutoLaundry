package repository

import (
	"github.com/NekruzRakhimov/AutoLaundery/db"
	"github.com/NekruzRakhimov/AutoLaundery/logger"
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/utils"
)

func GetAllServices() (services []models.Service, err error) {
	if err = db.GetDBConn().Table("services").Where("is_removed = false").Order("id").
		Scan(&services).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return nil, err
	}

	return
}

func GetServiceByCode(code string) (s models.Service, err error) {
	if err = db.GetDBConn().Table("services").Where("code = ?", code).
		Scan(&s).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return models.Service{}, err
	}

	return
}

func GetAllDiscounts() (d []models.Discount, err error) {
	if err = db.GetDBConn().Table("discounts").Where("is_removed = false").Order("id").
		Scan(&d).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return nil, err
	}

	return
}

func GetAllMarkUps() (d []models.MarkUp, err error) {
	if err = db.GetDBConn().Table("additional_markups").Where("is_removed = false").Order("id").
		Scan(&d).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return nil, err
	}

	return
}

func GetDiscountByID(id int) (d models.Discount, err error) {
	if err = db.GetDBConn().Table("discounts").Where("id = ?", id).
		Scan(&d).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return models.Discount{}, err
	}

	return
}

func GetMarkUpByIDs(id int) (d models.MarkUp, err error) {
	if err = db.GetDBConn().Table("additional_markups").Where("id = ?", id).
		Scan(&d).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return models.MarkUp{}, err
	}

	return
}
