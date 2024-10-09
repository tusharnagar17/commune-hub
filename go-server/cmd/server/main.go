package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tusharnagar17/commune-hub/internal/config"
	"github.com/tusharnagar17/commune-hub/internal/database"
	"github.com/tusharnagar17/commune-hub/internal/routes"
)

func main(){
	cfg := config.LoadConfig()
	// Initialize the database
	database.InitDatabase(cfg)

	// handle routes
	router := gin.Default()

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err!=nil {
		log.Fatal("Failed to run the server:", err)
	}
}