package service

import (
	"errors"
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/repository"
)

func GetAllServices() (s []models.Service, err error) {
	return repository.GetAllServices()
}

func GetServiceByCode(code string) (s models.Service, err error) {
	return repository.GetServiceByCode(code)
}

func GetServicePricing(code string) (sp []models.ServicePricing, err error) {
	switch code {
	case models.DryCleaning:
		return GetDryCleaningPricing()
	case models.HandWash:
		return GetHandWashPricing()
	case models.GeneralLaundryService:
		return GetGeneralLaundryService()
	case models.IroningServices:
		return GetIroningServicesPricing()
	case models.ClothingRepair:
		return GetClothingRepairPricing()
	case models.StainRemoval:
		return GetStainRemovalPricing()
	default:
		return nil, errors.New("no service with such code")
	}
}

func GetDryCleaningPricing() (sp []models.ServicePricing, err error) {
	return repository.GetDryCleaningPricing()
}

func GetHandWashPricing() (sp []models.ServicePricing, err error) {
	s, err := repository.GetServiceByCode(models.HandWash)
	if err != nil {
		return nil, err
	}

	sp = append(sp, models.ServicePricing{
		Name:  s.UniversalPriceTitle,
		Price: s.Price,
	})

	return
}

func GetGeneralLaundryService() (sp []models.ServicePricing, err error) {
	return repository.GetGeneralLaundryService()
}

func GetIroningServicesPricing() (sp []models.ServicePricing, err error) {
	return repository.GetIroningServicesPricing()
}

func GetClothingRepairPricing() (sp []models.ServicePricing, err error) {
	s, err := repository.GetServiceByCode(models.ClothingRepair)
	if err != nil {
		return nil, err
	}

	sp = append(sp, models.ServicePricing{
		Name:  s.UniversalPriceTitle,
		Price: s.Price,
	})

	return
}

func GetStainRemovalPricing() (sp []models.ServicePricing, err error) {
	return repository.GetStainRemovalPricing()
}
