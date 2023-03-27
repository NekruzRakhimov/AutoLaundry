package main

import (
	"github.com/NekruzRakhimov/AutoLaundery/db"
	"github.com/NekruzRakhimov/AutoLaundery/logger"
	"github.com/NekruzRakhimov/AutoLaundery/routes"
	"github.com/NekruzRakhimov/AutoLaundery/utils"
)

// @title Gin Swagger AutoLaundry Api
// @version 1.0
// @description Сервис для автоматизации работы химчистки.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email nekruzrakhimov@icloud.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8181
// @BasePath /
// @schemes http

func main() {
	utils.ReadSettings()
	utils.PutAdditionalSettings()
	logger.Init()
	db.StartDbConnection()
	//go jobs.StartJobs()
	routes.RunAllRoutes()
}
