package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	Date string `json:"date"`
	Location string `json:"host_id"`
	HostID uint `json:"host_id"`
	Host User `json:"host" gorm:"foreignKey:HostID"`
}