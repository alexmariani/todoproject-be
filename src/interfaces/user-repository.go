package interfaces

import (
	"todoproject-be/src/models"
)

type UserRepository interface {
	AddUser(body *models.User) error
	GetUser(username string) (*models.User, error)
	ResetPassword(username string, password string) error
}
