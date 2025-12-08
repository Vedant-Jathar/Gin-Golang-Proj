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

func (u *UserService) GetUserById(id int) (any, error) {

	var user models.User

	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) CreateUser(user *models.User) int {
	result := u.db.Create(user)
	if result.Error != nil {
		log.Fatal("Error creating the user in the service")
	}

	return user.Id
}

func (u *UserService) UpdateUser(user *models.User, id int) error {

	if err := u.db.First(&user, id).Error; err != nil {
		return err // record not found OR deleted
	}

	result := u.db.Model(&models.User{}).
		Where("Id = ?", id).
		Updates(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *UserService) DeleteUser(id int) error {

	result := u.db.Delete(&models.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil

}
