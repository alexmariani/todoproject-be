package repository

import (
	"todoproject-be/src/database"
	"todoproject-be/src/models"

	"golang.org/x/crypto/bcrypt"
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

func (ur UserRepository) ResetPassword(username string, password string) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	errSql := database.Db.Where(&models.User{Username: username}).Update("username,password", &models.User{Username: username, Password: string(hashedPassword)}).Error
	return errSql
}
