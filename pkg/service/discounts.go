package service

import (
	"github.com/NekruzRakhimov/AutoLaundery/logger"
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/repository"
)

func GetAllDiscounts() (d []models.Discount, err error) {
	return repository.GetAllDiscounts()
}

func GetBulkDiscountsByIDs(ids []int) (discounts []models.Discount, err error) {
	for _, id := range ids {
		discount, err := repository.GetDiscountByID(id)
		if err != nil {
			return nil, err
		}

		discounts = append(discounts, discount)
	}

	logger.Debug.Println(discounts)

	return
}

func GetTotalDiscountPercent(discounts []models.Discount) (totalDiscount float32) {
	for _, discount := range discounts {
		totalDiscount += discount.Percent
	}

	if totalDiscount > 100 {
		totalDiscount = 100
	}

	return
}

func GetTotalDiscountAmount(ids []int, amount float32) (discountAmount float32, err error) {
	discounts, err := GetBulkDiscountsByIDs(ids)
	if err != nil {
		return 0, err
	}

	totalDiscountPercent := GetTotalDiscountPercent(discounts)
	discountAmount = amount * (totalDiscountPercent / 100)
	return
}
