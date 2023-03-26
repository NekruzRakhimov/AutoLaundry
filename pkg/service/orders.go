package service

import (
	"github.com/NekruzRakhimov/AutoLaundery/db"
	"github.com/NekruzRakhimov/AutoLaundery/logger"
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/repository"
)

func MakeOrder(order models.Order) (id int, err error) {
	trx := db.GetDBConn().Begin()
	if err = repository.SaveOrder(trx, &order); err != nil {
		return 0, err
	}

	if order.DryCleaning != nil {
		totalAmount, err := CalculateDryCleaningOrder(order.DryCleaning)
		if err != nil {
			return 0, err
		}

		if err = repository.SaveDryCleaningOrder(trx, order.ID, order.DryCleaning); err != nil {
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

		if err = repository.SaveHandWashOrder(trx, order.ID, order.HandWash); err != nil {
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

		if err = repository.SaveGeneralLaundryServiceOrder(trx, order.ID, order.GeneralLaundryService); err != nil {
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

		if err = repository.SaveIroningServicesOrder(trx, order.ID, order.IroningServices); err != nil {
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

		if err = repository.SaveClothingRepairOrder(trx, order.ID, order.ClothingRepair); err != nil {
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

		if err = repository.SaveStainRemovalOrder(trx, order.ID, order.StainRemoval); err != nil {
			return 0, err
		}

		logger.Debug.Println("StainRemoval: ", totalAmount)
		order.TotalAmount += totalAmount
	}
	if err = repository.UpdateOrder(trx, &order); err != nil {
		return 0, err
	}
	trx.Commit()

	logger.Debug.Printf("%#v", order.TotalAmount)
	return order.ID, nil
}

func GetAllOrders() (orders []models.Order, err error) {
	return repository.GetAllOrders()
}

func GetOrderByID(id int) (order models.Order, err error) {
	return repository.GetOrderByID(id)
}
