package dto

type (
	QuizResponse struct {
		ID            string   `json:"id"`
		Question      string   `json:"question"`
		CorrectAnswer string   `json:"correct_answer"`
		WrongAnswer   []string `json:"wrong_answers"`
	}
)
