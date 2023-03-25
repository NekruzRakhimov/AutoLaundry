package controller

import (
	"github.com/NekruzRakhimov/AutoLaundery/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllService(c *gin.Context) {
	s, err := service.GetAllServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "ошибка при получении списка сервисов"})
	}

	if err != nil {
		return
	}
	c.JSON(http.StatusOK, s)
}

func GetServicePricing(c *gin.Context) {
	code := c.Param("code")

	sp, err := service.GetServicePricing(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "ошибка при получении списка сервисов"})
	}

	c.JSON(http.StatusOK, sp)
}
