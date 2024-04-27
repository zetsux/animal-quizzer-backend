package dto

type (
	QuestResponse struct {
		ID              string             `json:"id"`
		Step            int                `json:"step"`
		Hint            string             `json:"hint"`
		SilhouetteImage string             `json:"silhouette_image"`
		AnimalType      AnimalTypeResponse `json:"animal_type,omitempty"`
	}

	QuestLeaderboard struct {
		Username string
		Point    int
		Time     string
	}
)
