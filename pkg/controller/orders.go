package controller

import (
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MakeOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "некорректное тело запроса"})
		return
	}

	id, err := service.MakeOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "что-то пошло не так"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
