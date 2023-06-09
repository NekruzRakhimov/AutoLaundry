package routes

import (
	"fmt"
	"github.com/NekruzRakhimov/AutoLaundery/docs"
	"github.com/NekruzRakhimov/AutoLaundery/pkg/controller"
	"github.com/NekruzRakhimov/AutoLaundery/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

func RunAllRoutes() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	//r.Use(cors.Default())

	r.Use(controller.CORSMiddleware())

	// Статус код 500, при любых panic()
	r.Use(gin.Recovery())

	// Запуск роутов
	initAllRoutes(r)

	// Запуск сервера
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = utils.AppSettings.AppParams.PortRun
	}
	_ = r.Run(fmt.Sprintf(":%s", port))
}

func initAllRoutes(r *gin.Engine) {
	r.GET("/", controller.PingPong)
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/services", controller.GetAllService)
	r.GET("/services/:code/pricing", controller.GetServicePricing)

	r.POST("/order", controller.MakeOrder)
	r.GET("/order", controller.GetAllOrders)
	r.GET("/order/:id/details", controller.GetOrderByID)
	r.GET("/discounts", controller.GetAllDiscounts)
	r.GET("/mark_ups", controller.GetAllMarkUps)

}
