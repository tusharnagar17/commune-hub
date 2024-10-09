package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tusharnagar17/commune-hub/internal/database"
	"github.com/tusharnagar17/commune-hub/internal/models"
)

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err:= c.ShouldBindBodyWithJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error:": err.Error()})
		return
	}
	database.DB.Create(&event)
	c.JSON(http.StatusCreated, event)
}

func GetEvents(c *gin.Context) {
	var events []models.Event

	database.DB.Find(&events)
	c.JSON(http.StatusOK, events)
}
