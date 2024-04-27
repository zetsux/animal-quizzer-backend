package dto

type (
	QuizResponse struct {
		ID            string   `json:"id,omitempty"`
		Question      string   `json:"question,omitempty"`
		CorrectAnswer string   `json:"correct_answer,omitempty"`
		WrongAnswer   []string `json:"wrong_answers,omitempty"`
		Cooldown      int      `json:"cooldown,omitempty"`
	}
)
