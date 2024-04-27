package errors

import "errors"

var (
	ErrQuizCountNotEnough = errors.New("jumlah pertanyaan yang tersedia tidak cukup")
	ErrQuizInCooldown     = errors.New("kuis sedang dalam masa tunggu")
)
