package entity

import (
	"github.com/google/uuid"
	"github.com/zetsux/gin-gorm-clean-starter/common/base"
)

type Animal struct {
	ID                 uuid.UUID   `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name               string      `json:"name" gorm:"not null"`
	Step               int         `json:"step" gorm:"not null"`
	Hint               string      `json:"hint" gorm:"not null"`
	Description        string      `json:"description" gorm:"not null"`
	Habitat            string      `json:"habitat" gorm:"not null"`
	Food               string      `json:"food" gorm:"not null"`
	Reproduction       string      `json:"reproduction" gorm:"not null"`
	FunFact            string      `json:"fun_fact" gorm:"not null"`
	ConservationStatus string      `json:"conservation_status" gorm:"not null"`
	AnimalTypeID       string      `json:"animal_type_id" gorm:"foreignKey:AnimalTypeID"`
	AnimalType         *AnimalType `json:"animal_type,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SilhouetteImage    string      `json:"silhouette_image"`
	RealImage          string      `json:"real_image"`
	BadgeImage         string      `json:"badge_image"`
	base.Model
}
