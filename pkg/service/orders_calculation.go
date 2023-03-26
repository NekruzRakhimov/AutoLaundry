package service

import (
	"github.com/NekruzRakhimov/AutoLaundery/logger"
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/repository"
	"gorm.io/gorm"
)

func CalculateServicesTotalAmount(order models.Order, trx *gorm.DB) (totalAmount float32, err error) {
	if order.DryCleaning != nil {
		amount, err := CalculateDryCleaningOrder(order.DryCleaning)
		if err != nil {
			return 0, err
		}

		if err = repository.SaveDryCleaningOrder(trx, order.ID, order.DryCleaning); err != nil {
			return 0, err
		}

		logger.Debug.Println("DryCleaning: ", amount)
		totalAmount += amount
	}

	if order.HandWash != nil {
		amount, err := CalculateHandWashOrder(order.HandWash)
		if err != nil {
			return 0, err
		}

		if err = repository.SaveHandWashOrder(trx, order.ID, order.HandWash); err != nil {
			return 0, err
		}

		logger.Debug.Println("HandWash: ", amount)
		totalAmount += amount
	}

	if order.GeneralLaundryService != nil {
		amount, err := CalculateGeneralLaundryServiceOrder(order.GeneralLaundryService)
		if err != nil {
			return 0, err
		}

		if err = repository.SaveGeneralLaundryServiceOrder(trx, order.ID, order.GeneralLaundryService); err != nil {
			return 0, err
		}

		logger.Debug.Println("GeneralLaundryService: ", amount)
		totalAmount += amount
	}

	if order.IroningServices != nil {
		amount, err := CalculateIroningServicesOrder(order.IroningServices)
		if err != nil {
			return 0, err
		}

		if err = repository.SaveIroningServicesOrder(trx, order.ID, order.IroningServices); err != nil {
			return 0, err
		}

		logger.Debug.Println("IroningServices: ", amount)
		totalAmount += amount
	}

	if order.ClothingRepair != nil {
		amount, err := CalculateClothingRepairOrder(order.ClothingRepair)
		if err != nil {
			return 0, err
		}

		if err = repository.SaveClothingRepairOrder(trx, order.ID, order.ClothingRepair); err != nil {
			return 0, err
		}

		logger.Debug.Println("ClothingRepair: ", amount)
		totalAmount += amount
	}

	if order.StainRemoval != nil {
		amount, err := CalculateStainRemovalOrder(order.StainRemoval)
		if err != nil {
			return 0, err
		}

		if err = repository.SaveStainRemovalOrder(trx, order.ID, order.StainRemoval); err != nil {
			return 0, err
		}

		logger.Debug.Println("StainRemoval: ", amount)
		totalAmount += amount
	}

	return totalAmount, err
}

func CalculateDryCleaningOrder(order *[]models.DryCleaningOrder) (totalAmount float32, err error) {
	for i, cleaningOrder := range *order {
		p, err := repository.GetDryCleaningProductByID(cleaningOrder.ProductID)
		if err != nil {
			return 0, err
		}

		(*order)[i].Price = p.Price
		(*order)[i].TotalAmount = p.Price * float32(cleaningOrder.Quantity)

		var (
			discountAmount float32
			markUpAmount   float32
		)

		if cleaningOrder.DiscountIDs != nil {
			discountAmount, err = GetTotalDiscountAmount(*cleaningOrder.DiscountIDs, (*order)[i].TotalAmount)
			if err != nil {
				return 0, err
			}
		}
		(*order)[i].TotalAmount -= discountAmount

		if cleaningOrder.MarkUpIDs != nil {
			markUpAmount, err = GetTotalMarkUpAmount(*cleaningOrder.MarkUpIDs, (*order)[i].TotalAmount)
			if err != nil {
				return 0, err
			}
		}
		(*order)[i].TotalAmount += markUpAmount
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

	var (
		discountAmount float32
		markUpAmount   float32
	)

	if order.DiscountIDs != nil {
		discountAmount, err = GetTotalDiscountAmount(*order.DiscountIDs, order.TotalAmount)
		if err != nil {
			return 0, err
		}
	}
	order.TotalAmount -= discountAmount

	if order.MarkUpIDs != nil {
		markUpAmount, err = GetTotalMarkUpAmount(*order.MarkUpIDs, order.TotalAmount)
		if err != nil {
			return 0, err
		}
	}
	order.TotalAmount += markUpAmount
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

		var (
			discountAmount float32
			markUpAmount   float32
		)

		if cleaningOrder.DiscountIDs != nil {
			discountAmount, err = GetTotalDiscountAmount(*cleaningOrder.DiscountIDs, (*order)[i].TotalAmount)
			if err != nil {
				return 0, err
			}
		}
		(*order)[i].TotalAmount -= discountAmount

		if cleaningOrder.MarkUpIDs != nil {
			markUpAmount, err = GetTotalMarkUpAmount(*cleaningOrder.MarkUpIDs, (*order)[i].TotalAmount)
			if err != nil {
				return 0, err
			}
		}
		(*order)[i].TotalAmount += markUpAmount

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

		var (
			discountAmount float32
			markUpAmount   float32
		)

		if cleaningOrder.DiscountIDs != nil {
			discountAmount, err = GetTotalDiscountAmount(*cleaningOrder.DiscountIDs, (*order)[i].TotalAmount)
			if err != nil {
				return 0, err
			}
		}
		(*order)[i].TotalAmount -= discountAmount

		if cleaningOrder.MarkUpIDs != nil {
			markUpAmount, err = GetTotalMarkUpAmount(*cleaningOrder.MarkUpIDs, (*order)[i].TotalAmount)
			if err != nil {
				return 0, err
			}
		}
		(*order)[i].TotalAmount += markUpAmount
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

	var (
		discountAmount float32
		markUpAmount   float32
	)

	if order.DiscountIDs != nil {
		discountAmount, err = GetTotalDiscountAmount(*order.DiscountIDs, order.TotalAmount)
		if err != nil {
			return 0, err
		}
	}
	order.TotalAmount -= discountAmount

	if order.MarkUpIDs != nil {
		markUpAmount, err = GetTotalMarkUpAmount(*order.MarkUpIDs, order.TotalAmount)
		if err != nil {
			return 0, err
		}
	}
	order.TotalAmount += markUpAmount
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

		var (
			discountAmount float32
			markUpAmount   float32
		)

		if cleaningOrder.DiscountIDs != nil {
			discountAmount, err = GetTotalDiscountAmount(*cleaningOrder.DiscountIDs, (*order)[i].TotalAmount)
			if err != nil {
				return 0, err
			}
		}
		(*order)[i].TotalAmount -= discountAmount

		if cleaningOrder.MarkUpIDs != nil {
			markUpAmount, err = GetTotalMarkUpAmount(*cleaningOrder.MarkUpIDs, (*order)[i].TotalAmount)
			if err != nil {
				return 0, err
			}
		}
		(*order)[i].TotalAmount += markUpAmount
		totalAmount += (*order)[i].TotalAmount
	}

	return
}
