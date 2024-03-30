package repository

import (
	"todoproject-be/src/database"
	"todoproject-be/src/models"
)

type TipRepository struct{}

func (tr TipRepository) AddTip(body *models.Tip) error {
	err := database.Db.Create(body).Error
	return err
}

func (tr TipRepository) GetAllTips(idUtente int) ([]models.Tip, error) {
	var tips []models.Tip
	err := database.Db.Where(&models.Tip{IdOwner: idUtente}).Find(&tips).Error
	return tips, err
}

func (tr TipRepository) UpdateTip(body *models.Tip) (models.Tip, error) {
	err := database.Db.Save(&body).Error
	return *body, err
}

func (tr TipRepository) GetTip(idTip int) (models.Tip, error) {
	tip := models.Tip{}
	err := database.Db.Where(&models.Tip{ID: idTip}).Find(&tip).Error
	return tip, err
}
