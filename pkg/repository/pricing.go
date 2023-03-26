package repository

import (
	"github.com/NekruzRakhimov/AutoLaundery/db"
	"github.com/NekruzRakhimov/AutoLaundery/models"
)

func GetDryCleaningPricing() (sp []models.ServicePricing, err error) {
	if err = db.GetDBConn().Table("dry_cleaning_pricing").
		Joins("join products on products.id = dry_cleaning_pricing.product_id").
		Select("products.id, products.name, dry_cleaning_pricing.price").Scan(&sp).Error; err != nil {
		return nil, err
	}

	return
}

func GetGeneralLaundryServicePricing() (sp []models.ServicePricing, err error) {
	if err = db.GetDBConn().Table("general_laundry_pricing").
		Select("id, name, price").Scan(&sp).Error; err != nil {
		return nil, err
	}
	return
}

func GetIroningServicesPricing() (sp []models.ServicePricing, err error) {
	if err = db.GetDBConn().Table("ironing_service_pricing").
		Joins("join products on products.id = ironing_service_pricing.product_id").
		Select("products.id, products.name, ironing_service_pricing.price").Scan(&sp).Error; err != nil {
		return nil, err
	}

	return
}

func GetStainRemovalPricing() (sp []models.ServicePricing, err error) {
	if err = db.GetDBConn().Table("stain_removal_pricing").
		Select("id, name, price").Scan(&sp).Error; err != nil {
		return nil, err
	}
	return
}
