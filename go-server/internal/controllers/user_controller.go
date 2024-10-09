package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tusharnagar17/commune-hub/internal/database"
	"github.com/tusharnagar17/commune-hub/internal/models"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
	//Login Logic

	c.JSON(http.StatusOK, gin.H{"message": "Login Successful"})
}