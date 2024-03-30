package interfaces

import (
	"todoproject-be/src/models"
)

type TipRepository interface {
	AddTip(body *models.Tip) error
	GetAllTips(idUtente int) ([]models.Tip, error)
	UpdateTip(body *models.Tip) (models.Tip, error)
	GetTip(idTip int) (models.Tip, error)
}
