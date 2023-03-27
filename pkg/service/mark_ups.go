package service

import (
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/repository"
)

func GetAllMarkUps() (d []models.MarkUp, err error) {
	return repository.GetAllMarkUps()
}

func GetBulkMarkUpsByIDs(ids []int) (discounts []models.MarkUp, err error) {
	for _, id := range ids {
		discount, err := repository.GetMarkUpByIDs(id)
		if err != nil {
			return nil, err
		}

		discounts = append(discounts, discount)
	}

	return
}

func GetTotalMarkUpPercent(markUps []models.MarkUp) (totalDiscount float32) {
	for _, markUp := range markUps {
		totalDiscount += markUp.Percent
	}

	if totalDiscount > 100 {
		totalDiscount = 100
	}

	return
}

func GetTotalMarkUpAmount(ids []int, amount float32) (markUpAmount float32, err error) {
	markUps, err := GetBulkMarkUpsByIDs(ids)
	if err != nil {
		return 0, err
	}

	totalMarkUpPercent := GetTotalMarkUpPercent(markUps)
	markUpAmount = amount * (totalMarkUpPercent / 100)
	return
}
