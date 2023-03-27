package service

import (
	"github.com/NekruzRakhimov/AutoLaundery/db"
	"github.com/NekruzRakhimov/AutoLaundery/logger"
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/repository"
	"github.com/NekruzRakhimov/AutoLaundery/utils"
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

func GetAllOrders(date string) (orders []models.Order, err error) {
	return repository.GetAllOrders(date)
}

func GetOrderByID(id int) (order models.Order, err error) {
	orderDB, err := repository.GetOrderByID(id)
	if err != nil {
		return models.Order{}, err
	}

	order = models.Order{
		ID:          orderDB.ID,
		TotalAmount: orderDB.TotalAmount,
	}

	if orderDB.DiscountIDs != "{}" && orderDB.DiscountIDs != "" {
		discountIDs, err := utils.StringToArray(orderDB.DiscountIDs)
		if err != nil {
			return models.Order{}, err
		}
		discounts, err := GetBulkDiscountsByIDs(discountIDs)
		if err != nil {
			return models.Order{}, err
		}
		order.Discounts = &discounts
	}

	if orderDB.MarkUpIDs != "{}" && orderDB.MarkUpIDs != "" {
		markUpIDs, err := utils.StringToArray(orderDB.MarkUpIDs)
		if err != nil {
			return models.Order{}, err
		}
		markUps, err := GetBulkMarkUpsByIDs(markUpIDs)
		if err != nil {
			return models.Order{}, err
		}
		order.MarkUps = &markUps

	}

	if err = GetOrderServices(&order); err != nil {
		return models.Order{}, err
	}

	return
}

func GetDryCleaningOrderByID(orderID int) (ordersP *[]models.DryCleaningOrder, err error) {
	var orders []models.DryCleaningOrder
	ordersDB, err := repository.GetDryCleaningOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	for _, orderDB := range *ordersDB {
		order := models.DryCleaningOrder{
			ID:          orderDB.ID,
			OrderID:     orderDB.OrderID,
			ProductID:   orderDB.ProductID,
			ProductName: orderDB.ProductName,
			Quantity:    orderDB.Quantity,
			Price:       orderDB.Price,
			TotalAmount: orderDB.TotalAmount,
			//DiscountIDs: nil,
			//MarkUpIDs:   nil,
			//Discounts:   nil,
			//MarkUps:     nil,
		}

		if orderDB.DiscountIDs != "{}" && orderDB.DiscountIDs != "" {
			discountIDs, err := utils.StringToArray(orderDB.DiscountIDs)
			if err != nil {
				return nil, err
			}
			discounts, err := GetBulkDiscountsByIDs(discountIDs)
			if err != nil {
				return nil, err
			}
			order.Discounts = &discounts
		}

		if orderDB.MarkUpIDs != "{}" && orderDB.MarkUpIDs != "" {
			markUpIDs, err := utils.StringToArray(orderDB.MarkUpIDs)
			if err != nil {
				return nil, err
			}
			markUps, err := GetBulkMarkUpsByIDs(markUpIDs)
			if err != nil {
				return nil, err
			}
			order.MarkUps = &markUps
		}
		orders = append(orders, order)
	}

	return &orders, nil
}

func GetGeneralLaundryServiceOrderByID(orderID int) (ordersP *[]models.GeneralLaundryServiceOrder, err error) {
	var orders []models.GeneralLaundryServiceOrder
	ordersDB, err := repository.GetGeneralLaundryServiceOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	for _, orderDB := range *ordersDB {
		order := models.GeneralLaundryServiceOrder{
			ID:              orderDB.ID,
			OrderID:         orderDB.OrderID,
			ProductTypeID:   orderDB.ProductTypeID,
			ProductTypeName: orderDB.ProductTypeName,
			Weight:          orderDB.Weight,
			Price:           orderDB.Price,
			TotalAmount:     orderDB.TotalAmount,
		}

		if orderDB.DiscountIDs != "{}" && orderDB.DiscountIDs != "" {
			discountIDs, err := utils.StringToArray(orderDB.DiscountIDs)
			if err != nil {
				return nil, err
			}
			discounts, err := GetBulkDiscountsByIDs(discountIDs)
			if err != nil {
				return nil, err
			}
			order.Discounts = &discounts
		}

		if orderDB.MarkUpIDs != "{}" && orderDB.MarkUpIDs != "" {
			markUpIDs, err := utils.StringToArray(orderDB.MarkUpIDs)
			if err != nil {
				return nil, err
			}
			markUps, err := GetBulkMarkUpsByIDs(markUpIDs)
			if err != nil {
				return nil, err
			}
			order.MarkUps = &markUps
		}
		orders = append(orders, order)
	}

	return &orders, nil
}

func GetHandWashOrderByID(orderID int) (order *models.HandWashOrder, err error) {
	orderDB, err := repository.GetHandWashOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	order = &models.HandWashOrder{
		ID:          orderDB.ID,
		OrderID:     orderDB.OrderID,
		Quantity:    orderDB.Quantity,
		Price:       orderDB.Price,
		TotalAmount: orderDB.TotalAmount,
	}

	if orderDB.DiscountIDs != "{}" && orderDB.DiscountIDs != "" {
		discountIDs, err := utils.StringToArray(orderDB.DiscountIDs)
		if err != nil {
			return nil, err
		}
		discounts, err := GetBulkDiscountsByIDs(discountIDs)
		if err != nil {
			return nil, err
		}
		order.Discounts = &discounts
	}

	if orderDB.MarkUpIDs != "{}" && orderDB.MarkUpIDs != "" {
		markUpIDs, err := utils.StringToArray(orderDB.MarkUpIDs)
		if err != nil {
			return nil, err
		}
		markUps, err := GetBulkMarkUpsByIDs(markUpIDs)
		if err != nil {
			return nil, err
		}
		order.MarkUps = &markUps
	}

	return
}

func GetIroningServicesOrderByID(orderID int) (ordersP *[]models.IroningServicesOrder, err error) {
	var orders []models.IroningServicesOrder
	ordersDB, err := repository.GetIroningServicesOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	for _, orderDB := range *ordersDB {
		order := models.IroningServicesOrder{
			ID:          orderDB.ID,
			OrderID:     orderDB.OrderID,
			ProductID:   orderDB.ProductID,
			ProductName: orderDB.ProductName,
			Quantity:    orderDB.Quantity,
			Price:       orderDB.Price,
			TotalAmount: orderDB.TotalAmount,
		}

		if orderDB.DiscountIDs != "{}" && orderDB.DiscountIDs != "" {
			discountIDs, err := utils.StringToArray(orderDB.DiscountIDs)
			if err != nil {
				return nil, err
			}
			discounts, err := GetBulkDiscountsByIDs(discountIDs)
			if err != nil {
				return nil, err
			}
			order.Discounts = &discounts
		}

		if orderDB.MarkUpIDs != "{}" && orderDB.MarkUpIDs != "" {
			markUpIDs, err := utils.StringToArray(orderDB.MarkUpIDs)
			if err != nil {
				return nil, err
			}
			markUps, err := GetBulkMarkUpsByIDs(markUpIDs)
			if err != nil {
				return nil, err
			}
			order.MarkUps = &markUps
		}
		orders = append(orders, order)
	}

	return &orders, nil
}

func GetClothingRepairOrderByID(orderID int) (order *models.ClothingRepairOrder, err error) {
	orderDB, err := repository.GetHandWashOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	order = &models.ClothingRepairOrder{
		ID:          orderDB.ID,
		OrderID:     orderDB.OrderID,
		Quantity:    orderDB.Quantity,
		Price:       orderDB.Price,
		TotalAmount: orderDB.TotalAmount,
	}

	if orderDB.DiscountIDs != "{}" && orderDB.DiscountIDs != "" {
		discountIDs, err := utils.StringToArray(orderDB.DiscountIDs)
		if err != nil {
			return nil, err
		}
		discounts, err := GetBulkDiscountsByIDs(discountIDs)
		if err != nil {
			return nil, err
		}
		order.Discounts = &discounts
	}

	if orderDB.MarkUpIDs != "{}" && orderDB.MarkUpIDs != "" {
		markUpIDs, err := utils.StringToArray(orderDB.MarkUpIDs)
		if err != nil {
			return nil, err
		}
		markUps, err := GetBulkMarkUpsByIDs(markUpIDs)
		if err != nil {
			return nil, err
		}
		order.MarkUps = &markUps
	}

	return
}

func GetStainRemovalOrderByID(orderID int) (ordersP *[]models.StainRemovalOrder, err error) {
	var orders []models.StainRemovalOrder
	ordersDB, err := repository.GetStainRemovalOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	for _, orderDB := range *ordersDB {
		order := models.StainRemovalOrder{
			ID:          orderDB.ID,
			OrderID:     orderDB.OrderID,
			TypeID:      orderDB.TypeID,
			TypeName:    orderDB.TypeName,
			Quantity:    orderDB.Quantity,
			Price:       orderDB.Price,
			TotalAmount: orderDB.TotalAmount,
		}

		if orderDB.DiscountIDs != "{}" && orderDB.DiscountIDs != "" {
			discountIDs, err := utils.StringToArray(orderDB.DiscountIDs)
			if err != nil {
				return nil, err
			}
			discounts, err := GetBulkDiscountsByIDs(discountIDs)
			if err != nil {
				return nil, err
			}
			order.Discounts = &discounts
		}

		if orderDB.MarkUpIDs != "{}" && orderDB.MarkUpIDs != "" {
			markUpIDs, err := utils.StringToArray(orderDB.MarkUpIDs)
			if err != nil {
				return nil, err
			}
			markUps, err := GetBulkMarkUpsByIDs(markUpIDs)
			if err != nil {
				return nil, err
			}
			order.MarkUps = &markUps
		}
		orders = append(orders, order)
	}

	return &orders, nil
}

func GetOrderServices(order *models.Order) (err error) {
	order.DryCleaning, err = GetDryCleaningOrderByID(order.ID)
	if err != nil {
		return err
	}

	order.HandWash, err = GetHandWashOrderByID(order.ID)
	if err != nil {
		return err
	}

	order.GeneralLaundryService, err = GetGeneralLaundryServiceOrderByID(order.ID)
	if err != nil {
		return err
	}

	order.IroningServices, err = GetIroningServicesOrderByID(order.ID)
	if err != nil {
		return err
	}

	order.ClothingRepair, err = GetClothingRepairOrderByID(order.ID)
	if err != nil {
		return err
	}

	order.StainRemoval, err = GetStainRemovalOrderByID(order.ID)
	if err != nil {
		return err
	}

	return nil
}
