package controller

import (
	"github.com/NekruzRakhimov/AutoLaundery/logger"
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/service"
	"github.com/NekruzRakhimov/AutoLaundery/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func MakeOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "некорректное тело запроса"})
		return
	}

	id, totalAmount, err := service.MakeOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "что-то пошло не так"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order_id":     id,
		"total_amount": totalAmount,
	})
}

func GetAllOrders(c *gin.Context) {
	date := c.Query("date")

	if date != "" {
		if err := utils.ValidateDateFormat(date); err != nil {
			logger.Error.Printf("[%] Error is: %s", utils.FuncName(), "error while parsing 'date' query param")
			c.JSON(http.StatusBadRequest, gin.H{"reason": "дата должна быть в формате: 'YYYY-MM-DD'"})
			return
		}
	}

	o, err := service.GetAllOrders(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "что-то пошло не так"})
		return
	}

	c.JSON(http.StatusOK, o)
}

func GetOrderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[%] Error is: %s", utils.FuncName(), err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "некорректный номер заказа"})
		return
	}

	o, err := service.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "что-то пошло не так"})
		return
	}

	c.JSON(http.StatusOK, o)
}
