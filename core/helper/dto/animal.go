package dto

type (
	AnimalResponse struct {
		ID                 string             `json:"id,omitempty"`
		Name               string             `json:"name,omitempty"`
		Step               int                `json:"step,omitempty"`
		Hint               string             `json:"hint,omitempty"`
		Description        string             `json:"description,omitempty"`
		Habitat            string             `json:"habitat,omitempty"`
		Food               string             `json:"food,omitempty"`
		Reproduction       string             `json:"reproduction,omitempty"`
		FunFact            string             `json:"fun_fact,omitempty"`
		IsOwned            *bool              `json:"is_owned,omitempty"`
		ConservationStatus string             `json:"conservation_status,omitempty"`
		AnimalType         AnimalTypeResponse `json:"animal_type,omitempty"`
		SilhouetteImage    string             `json:"silhouette_image,omitempty"`
		RealImage          string             `json:"real_image,omitempty"`
		BadgeImage         string             `json:"badge_image,omitempty"`
	}
)
