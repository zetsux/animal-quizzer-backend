package entity

import (
	"github.com/google/uuid"
	"github.com/zetsux/gin-gorm-clean-starter/common/base"
)

type Animal struct {
	ID           uuid.UUID   `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name         string      `gorm:"name" json:"name"`
	Step         int         `gorm:"step" json:"step"`
	AnimalTypeID string      `json:"animal_type_id" gorm:"foreignKey:AnimalTypeID"`
	AnimalType   *AnimalType `json:"animal_type,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	base.Model
}
