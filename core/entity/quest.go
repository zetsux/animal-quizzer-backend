package entity

import (
	"github.com/google/uuid"
	"github.com/zetsux/gin-gorm-clean-starter/common/base"
)

type Quest struct {
	ID           uuid.UUID   `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Step         int         `gorm:"step" json:"step"`
	UserID       string      `json:"user_id" gorm:"foreignKey:UserID"`
	User         *User       `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AnimalTypeID string      `json:"animal_type_id" gorm:"foreignKey:AnimalTypeID"`
	AnimalType   *AnimalType `json:"animal_type,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	base.Model
}
