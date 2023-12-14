package controllers

import (
	"github.com/vikash-parashar/task-remainder/config"
	"github.com/vikash-parashar/task-remainder/models"
)

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
	err := config.DB.Create(user).Error
	return err
}

// GetUserByUsernameOrEmail retrieves a user by username or email
func GetUserByUsernameOrEmail(identifier string) (*models.User, error) {
	user := &models.User{}
	err := config.DB.Where("username = ? OR email = ?", identifier, identifier).First(user).Error
	return user, err
}
