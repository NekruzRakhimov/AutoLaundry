package controller

import (
	"github.com/NekruzRakhimov/AutoLaundery/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllService AutoLaundry godoc
// @Summary AutoLaundry
// @Description Роут для получения списка сервисов
// @Accept  json
// @Produce  json
// @Tags services
// @Success 200 {array} models.Service
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /services [get]
func GetAllService(c *gin.Context) {
	s, err := service.GetAllServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "ошибка при получении списка сервисов"})
		return
	}

	c.JSON(http.StatusOK, s)
}

// GetServicePricing AutoLaundry godoc
// @Summary AutoLaundry
// @Description Роут для получения информации о расценке каждого сервиса по коду
// @Accept json
// @Produce json
// @Tags services
// @Param code path string true "Код сервиса"
// @Success 200 {array} models.ServicePricing
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /services/{code}/pricing [get]
func GetServicePricing(c *gin.Context) {
	code := c.Param("code")

	sp, err := service.GetServicePricing(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "ошибка при получении списка сервисов"})
		return
	}

	c.JSON(http.StatusOK, sp)
}

// GetAllDiscounts AutoLaundry godoc
// @Summary AutoLaundry
// @Description Роут для получения списка скидок
// @Accept json
// @Produce json
// @Tags additional
// @Success 200 {array} models.Discount
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /discounts [get]
func GetAllDiscounts(c *gin.Context) {
	d, err := service.GetAllDiscounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "ошибка при получении списка скидок"})
		return
	}

	c.JSON(http.StatusOK, d)
}

// GetAllMarkUps AutoLaundry godoc
// @Summary AutoLaundry
// @Description Роут для получения списка наценок
// @Accept json
// @Produce json
// @Tags additional
// @Success 200 {array} models.MarkUp
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /mark_ups [get]
func GetAllMarkUps(c *gin.Context) {
	d, err := service.GetAllMarkUps()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "ошибка при получении списка наценок"})
		return
	}

	c.JSON(http.StatusOK, d)
}
