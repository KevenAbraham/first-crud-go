package main

import (
	"log"

	"github.com/KevenAbraham/first-crud-go/src/configurations/logger"
	"github.com/KevenAbraham/first-crud-go/src/handler/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start the application")
	err := godotenv.Load("variables.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
