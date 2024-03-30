package repository

import (
	"todoproject-be/src/database"
	"todoproject-be/src/models"
)

type UserRepository struct{}

func (ur UserRepository) AddUser(body *models.User) error {
	err := database.Db.Create(&body).Error
	return err
}

func (ur UserRepository) GetUser(username string) (*models.User, error) {
	user := &models.User{}
	err := database.Db.Where(&models.User{Username: username}).First(user).Error
	return user, err
}
