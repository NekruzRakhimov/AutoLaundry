package service

import (
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/repository"
)

func GetAllServices() (s []models.Service, err error) {
	return repository.GetAllServices()
}
