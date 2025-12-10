package services

import (
	// "github.com/Vedant-Jathar/Gin_Project/internal/models"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func (a *AuthService) InitAuthservice(database *gorm.DB) *AuthService {
	return &AuthService{
		db: database,
	}
}

// func (a *AuthService) Login(data models.AuthUser) (*models.AuthUser, error) {
	
// }
