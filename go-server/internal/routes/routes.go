package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tusharnagar17/commune-hub/internal/controllers"
	// "community-platform/controllers" // Update the path as per your project structure
	// "community-platform/middleware"
)

// SetupRoutes defines all the routes for the application
func SetupRoutes(router *gin.Engine) {
	// Test routes
	router.GET("/test", func(c *gin.Context){
		c.IndentedJSON(http.StatusOK, gin.H{"Status Check": "OK"})	
	})

	// User routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Protected routes
	authorized := router.Group("/")
	// authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/events", controllers.CreateEvent)
		authorized.GET("/events", controllers.GetEvents)
		// authorized.PUT("/posts/:id", controllers.UpdatePost)
		// authorized.DELETE("/posts/:id", controllers.DeletePost)
	}
}
