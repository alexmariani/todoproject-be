package models

type Tip struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Description string `json:"description" gorm:"not null"`
	IdOwner     int    `json:"idOwner" gorm:"foreignKey:UserID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
