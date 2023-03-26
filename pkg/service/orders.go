package service

import (
	"github.com/NekruzRakhimov/AutoLaundery/db"
	"github.com/NekruzRakhimov/AutoLaundery/logger"
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/repository"
)

func MakeOrder(order models.Order) (id int, totalAmount float32, err error) {
	trx := db.GetDBConn().Begin()
	if err = repository.SaveOrder(trx, &order); err != nil {
		return 0, 0, err
	}

	totalAmount, err = CalculateServicesTotalAmount(order, trx)
	if err != nil {
		return 0, 0, err
	}
	order.TotalAmount = totalAmount

	var (
		discountAmount float32
		markUpAmount   float32
	)

	if order.DiscountIDs != nil {
		discountAmount, err = GetTotalDiscountAmount(*order.DiscountIDs, order.TotalAmount)
		if err != nil {
			return 0, 0, err
		}
	}
	order.TotalAmount -= discountAmount

	if order.MarkUpIDs != nil {
		markUpAmount, err = GetTotalMarkUpAmount(*order.MarkUpIDs, order.TotalAmount)
		if err != nil {
			return 0, 0, err
		}
	}
	order.TotalAmount += markUpAmount

	if err = repository.UpdateOrder(trx, &order); err != nil {
		return 0, 0, err
	}
	trx.Commit()

	logger.Debug.Printf("%#v", order.TotalAmount)
	return order.ID, order.TotalAmount, nil
}

func GetAllOrders() (orders []models.Order, err error) {
	return repository.GetAllOrders()
}

func GetOrderByID(id int) (order models.Order, err error) {
	return repository.GetOrderByID(id)
}
