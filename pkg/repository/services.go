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

func GetDryCleaningPricing() (sp []models.ServicePricing, err error) {
	if err = db.GetDBConn().Table("dry_cleaning_pricing").
		Joins("join products on products.id = dry_cleaning_pricing.product_id").
		Select("products.name, dry_cleaning_pricing.price").Scan(&sp).Error; err != nil {
		return nil, err
	}

	return
}

func GetGeneralLaundryService() (sp []models.ServicePricing, err error) {
	if err = db.GetDBConn().Table("general_laundry_pricing").
		Select("name, price").Scan(&sp).Error; err != nil {
		return nil, err
	}
	return
}

func GetIroningServicesPricing() (sp []models.ServicePricing, err error) {
	if err = db.GetDBConn().Table("ironing_service_pricing").
		Joins("join products on products.id = ironing_service_pricing.product_id").
		Select("products.name, ironing_service_pricing.price").Scan(&sp).Error; err != nil {
		return nil, err
	}

	return
}

func GetStainRemovalPricing() (sp []models.ServicePricing, err error) {
	if err = db.GetDBConn().Table("stain_removal_pricing").
		Select("name, price").Scan(&sp).Error; err != nil {
		return nil, err
	}
	return
}
