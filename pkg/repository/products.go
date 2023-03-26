package repository

import (
	"fmt"
	"github.com/NekruzRakhimov/AutoLaundery/db"
	"github.com/NekruzRakhimov/AutoLaundery/logger"
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/utils"
)

func GetDryCleaningProductByID(id int) (p models.ServicePricing, err error) {
	if err = db.GetDBConn().Table("dry_cleaning_pricing").Select("price").
		Where("product_id = ?", id).Scan(&p).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return models.ServicePricing{}, err
	}

	return
}

func GetGeneralLaundryServiceProductByID(id int) (p models.ServicePricing, err error) {
	fmt.Println(id)
	if err = db.GetDBConn().Table("general_laundry_pricing").Select("price").
		Where("id = ?", id).Scan(&p).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return models.ServicePricing{}, err
	}

	return
}

func GetIroningServicesProductByID(id int) (p models.ServicePricing, err error) {
	if err = db.GetDBConn().Table("ironing_service_pricing").Select("price").
		Where("product_id = ?", id).Scan(&p).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return models.ServicePricing{}, err
	}

	return
}

func GetStainRemovalProductByID(id int) (p models.ServicePricing, err error) {
	if err = db.GetDBConn().Table("stain_removal_pricing").Select("price").
		Where("id = ?", id).Scan(&p).Error; err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		return models.ServicePricing{}, err
	}

	return
}
