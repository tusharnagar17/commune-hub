package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tusharnagar17/commune-hub/internal/controllers"
	"github.com/tusharnagar17/commune-hub/internal/middleware"
	// "community-platform/controllers" // Update the path as per your project structure
	// "community-platform/middleware"
)

// SetupRoutes defines all the routes for the application
func SetupRoutes(router *gin.Engine) {
	// Test routes
	router.GET("/test", func(c *gin.Context){
		log.Println("GET at /test")
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// User routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Protected routes
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{	
		
		authorized.POST("/events", controllers.CreateEvent)
		authorized.GET("/events", controllers.GetEvents)
		// authorized.PUT("/posts/:id", controllers.UpdatePost)
		// authorized.DELETE("/posts/:id", controllers.DeletePost)
	}
}
