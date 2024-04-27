package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/zetsux/gin-gorm-clean-starter/common/base"
)

type Quiz struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Question      string         `json:"question" gorm:"not null"`
	CorrectAnswer string         `json:"correct_answer" gorm:"not null"`
	WrongAnswer   pq.StringArray `gorm:"type:varchar(255)[];not null" json:"wrong_answers"`
	AnimalID      string         `json:"animal_id" gorm:"foreignKey:AnimalID"`
	Animal        *Animal        `json:"animal,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	base.Model
}
