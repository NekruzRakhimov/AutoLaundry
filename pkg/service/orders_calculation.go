package service

import (
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/repository"
)

func CalculateDryCleaningOrder(order *[]models.DryCleaningOrder) (totalAmount float32, err error) {
	for i, cleaningOrder := range *order {
		p, err := repository.GetDryCleaningProductByID(cleaningOrder.ProductID)
		if err != nil {
			return 0, err
		}

		(*order)[i].Price = p.Price
		(*order)[i].TotalAmount = p.Price * float32(cleaningOrder.Quantity)
		totalAmount += (*order)[i].TotalAmount
	}

	return
}

func CalculateHandWashOrder(order *models.HandWashOrder) (totalAmount float32, err error) {
	s, err := repository.GetServiceByCode(models.HandWashCode)
	if err != nil {
		return 0, err
	}

	order.Price = s.Price
	order.TotalAmount = s.Price * float32(order.Quantity)
	totalAmount += order.TotalAmount
	return
}

func CalculateGeneralLaundryServiceOrder(order *[]models.GeneralLaundryServiceOrder) (totalAmount float32, err error) {
	for i, cleaningOrder := range *order {
		p, err := repository.GetGeneralLaundryServiceProductByID(cleaningOrder.ProductTypeID)
		if err != nil {
			return 0, err
		}

		(*order)[i].Price = p.Price
		(*order)[i].TotalAmount = p.Price * cleaningOrder.Weight
		totalAmount += (*order)[i].TotalAmount
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
		(*order)[i].TotalAmount = p.Price * float32(cleaningOrder.Quantity)
		totalAmount += (*order)[i].TotalAmount
	}

	return
}

func CalculateClothingRepairOrder(order *models.ClothingRepairOrder) (totalAmount float32, err error) {
	s, err := repository.GetServiceByCode(models.ClothingRepairCode)
	if err != nil {
		return 0, err
	}

	order.Price = s.Price
	order.TotalAmount = s.Price * float32(order.Quantity)
	totalAmount += order.TotalAmount
	return
}

func CalculateStainRemovalOrder(order *[]models.StainRemovalOrder) (totalAmount float32, err error) {
	for i, cleaningOrder := range *order {
		p, err := repository.GetStainRemovalProductByID(cleaningOrder.TypeID)
		if err != nil {
			return 0, err
		}

		(*order)[i].Price = p.Price
		(*order)[i].TotalAmount = p.Price * float32(cleaningOrder.Quantity)
		totalAmount += (*order)[i].TotalAmount
	}

	return
}
