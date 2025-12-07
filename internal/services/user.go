package services

import (
	"github.com/Vedant-Jathar/Gin_Project/internal/models"
	"gorm.io/gorm"
	"log"
)

type UserService struct {
	db *gorm.DB
}

func (u *UserService) NewUserService(database *gorm.DB) *UserService {

	return &UserService{
		db: database,
	}
}

func (u *UserService) GetUsers() ([]models.User, error) {

	var users []models.User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *UserService) CreateUser(user *models.User) int {
	result := u.db.Create(user)
	if result.Error != nil {
		log.Fatal("Error creating the user in the service")
	}

	return user.Id
}
