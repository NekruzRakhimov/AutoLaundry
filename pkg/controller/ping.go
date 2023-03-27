package controller

import (
	"github.com/NekruzRakhimov/AutoLaundery/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingPong AutoLaundry godoc
// @Summary AutoLaundry
// @Description Роут для проверки работы сервера
// @Accept  json
// @Produce  json
// @Tags base
// @Success 200 {object} models.PingPong
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router / [get]
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, models.PingPong{Pong: "Server is up and running!"})
}
