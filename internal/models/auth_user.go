package models

import "gorm.io/gorm"

type AuthUser struct {
	gorm.Model
	Id       int    `gorm:"primaryKey"`
	Name     int    `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

