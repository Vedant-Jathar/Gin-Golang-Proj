package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id    int    `gorm:"primaryKey"`
	Email string `gorm:"uniqueIndex;not null" json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}
