package service

import (
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/repository"
)

func GetAllMarkUps() (d []models.MarkUp, err error) {
	return repository.GetAllMarkUps()
}

func GetBulkMarkUpsByIDs(ids []int) (discounts []models.Discount, err error) {
	for _, id := range ids {
		discount, err := repository.GetMarkUpByIDs(id)
		if err != nil {
			return nil, err
		}

		discounts = append(discounts, discount)
	}

	return
}

func GetTotalMarkUpPercent(discounts []models.Discount) (totalDiscount float32) {
	for _, discount := range discounts {
		totalDiscount += discount.Percent
	}

	if totalDiscount > 100 {
		totalDiscount = 100
	}

	return
}

func GetTotalMarkUpAmount(ids []int, amount float32) (markUpAmount float32, err error) {
	discounts, err := GetBulkMarkUpsByIDs(ids)
	if err != nil {
		return 0, err
	}

	totalMarkUpPercent := GetTotalMarkUpPercent(discounts)
	markUpAmount = amount * (totalMarkUpPercent / 100)
	return
}
