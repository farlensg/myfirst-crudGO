package main

import (
	"log"

	"github.com/farlensg/myfirst-crudGO/src/configuration/logger"
	"github.com/farlensg/myfirst-crudGO/src/controller"
	"github.com/farlensg/myfirst-crudGO/src/controller/routes"
	"github.com/farlensg/myfirst-crudGO/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Abount to start user application")

	// Substitua log.Fatal por um aviso de erro amigável
	if err := godotenv.Load(); err != nil {
		logger.Info("Aviso: arquivo .env não encontrado. Carregando variáveis de ambiente nativas.")
	}

	// Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// package main

// import (
// 	"log"

// 	"github.com/farlensg/myfirst-crudGO/src/configuration/logger"
// 	"github.com/farlensg/myfirst-crudGO/src/controller/routes"
// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	logger.Info("Abount to start user application")
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	router := gin.Default()

// 	routes.InitRoutes(&router.RouterGroup)

// 	if err := router.Run(":8080"); err != nil {
// 		log.Fatal(err)
// 	}
// }
