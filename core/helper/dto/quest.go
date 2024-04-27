package dto

type (
	QuestResponse struct {
		ID         string             `json:"id"`
		Step       int                `json:"step"`
		AnimalType AnimalTypeResponse `json:"animal_type,omitempty"`
	}
)
