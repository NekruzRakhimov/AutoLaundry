package service

import (
	"github.com/NekruzRakhimov/AutoLaundery/logger"
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/repository"
)

func MakeOrder(order models.Order) (id int, err error) {
	if order.DryCleaning != nil {
		totalAmount, err := CalculateDryCleaningOrder(order.DryCleaning)
		if err != nil {
			return 0, err
		}

		logger.Debug.Println("DryCleaning: ", totalAmount)
		order.TotalAmount += totalAmount
	}

	if order.HandWash != nil {
		totalAmount, err := CalculateHandWashOrder(order.HandWash)
		if err != nil {
			return 0, err
		}

		logger.Debug.Println("HandWash: ", totalAmount)
		order.TotalAmount += totalAmount
	}

	if order.GeneralLaundryService != nil {
		totalAmount, err := CalculateGeneralLaundryServiceOrder(order.GeneralLaundryService)
		if err != nil {
			return 0, err
		}

		logger.Debug.Println("GeneralLaundryService: ", totalAmount)
		order.TotalAmount += totalAmount
	}

	if order.IroningServices != nil {
		totalAmount, err := CalculateIroningServicesOrder(order.IroningServices)
		if err != nil {
			return 0, err
		}

		logger.Debug.Println("IroningServices: ", totalAmount)
		order.TotalAmount += totalAmount
	}

	if order.ClothingRepair != nil {
		totalAmount, err := CalculateClothingRepairOrder(order.ClothingRepair)
		if err != nil {
			return 0, err
		}

		logger.Debug.Println("ClothingRepair: ", totalAmount)
		order.TotalAmount += totalAmount
	}

	if order.StainRemoval != nil {
		totalAmount, err := CalculateStainRemovalOrder(order.StainRemoval)
		if err != nil {
			return 0, err
		}

		logger.Debug.Println("StainRemoval: ", totalAmount)
		order.TotalAmount += totalAmount
	}

	logger.Debug.Printf("%#v", order.TotalAmount)
	return 0, err
}

func CalculateDryCleaningOrder(order *[]models.DryCleaningOrder) (totalAmount float32, err error) {
	for i, cleaningOrder := range *order {
		p, err := repository.GetDryCleaningProductByID(cleaningOrder.ProductID)
		if err != nil {
			return 0, err
		}

		(*order)[i].Price = p.Price
		totalAmount += p.Price * float32(cleaningOrder.Quantity)
	}

	return
}

func CalculateHandWashOrder(order *models.HandWashOrder) (totalAmount float32, err error) {
	s, err := repository.GetServiceByCode(models.HandWashCode)
	if err != nil {
		return 0, err
	}

	totalAmount += s.Price * float32(order.Quantity)
	return
}

func CalculateGeneralLaundryServiceOrder(order *[]models.GeneralLaundryServiceOrder) (totalAmount float32, err error) {
	for i, cleaningOrder := range *order {
		p, err := repository.GetGeneralLaundryServiceProductByID(cleaningOrder.ProductTypeID)
		if err != nil {
			return 0, err
		}

		(*order)[i].Price = p.Price
		totalAmount += p.Price * cleaningOrder.Weight
	}

	return
}

func CalculateIroningServicesOrder(order *[]models.IroningServicesOrder) (totalAmount float32, err error) {
	for i, cleaningOrder := range *order {
		p, err := repository.GetIroningServicesProductByID(cleaningOrder.ProductID)
		if err != nil {
			return 0, err
		}

		(*order)[i].Price = p.Price
		totalAmount += p.Price * float32(cleaningOrder.Quantity)
	}

	return
}

func CalculateClothingRepairOrder(order *models.ClothingRepairOrder) (totalAmount float32, err error) {
	s, err := repository.GetServiceByCode(models.ClothingRepairCode)
	if err != nil {
		return 0, err
	}

	totalAmount += s.Price * float32(order.Quantity)
	return
}

func CalculateStainRemovalOrder(order *[]models.StainRemovalOrder) (totalAmount float32, err error) {
	for i, cleaningOrder := range *order {
		p, err := repository.GetStainRemovalProductByID(cleaningOrder.TypeID)
		if err != nil {
			return 0, err
		}

		(*order)[i].Price = p.Price
		totalAmount += p.Price * float32(cleaningOrder.Quantity)
	}

	return
}
