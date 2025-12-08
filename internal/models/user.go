package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id    int    `gorm:"primaryKey"`
	Email string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Name  string `gorm:"type:varchar(255)" json:"name" binding:"required"`
	Age   int    `json:"age"`
}