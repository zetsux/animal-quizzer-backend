package dto

type (
	AnimalResponse struct {
		ID              string             `json:"id"`
		Name            string             `json:"name"`
		Step            int                `json:"step"`
		Hint            string             `json:"hint,omitempty"`
		AnimalType      AnimalTypeResponse `json:"animal_type,omitempty"`
		SilhouetteImage string             `json:"silhouette_image,omitempty"`
		RealImage       string             `json:"real_image,omitempty"`
		BadgeImage      string             `json:"badge_image,omitempty"`
	}
)
