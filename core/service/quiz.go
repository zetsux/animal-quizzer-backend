package service

import (
	"context"
	"math/rand"

	"github.com/zetsux/gin-gorm-clean-starter/core/helper/dto"
	errs "github.com/zetsux/gin-gorm-clean-starter/core/helper/errors"
	"github.com/zetsux/gin-gorm-clean-starter/core/repository"
)

type quizService struct {
	quizRepository   repository.QuizRepository
	animalRepository repository.AnimalRepository
}

type QuizService interface {
	GetAnimalQuiz(ctx context.Context, userID string, animalID string) ([]dto.QuizResponse, error)
}

func NewQuizService(quizR repository.QuizRepository, animalR repository.AnimalRepository) QuizService {
	return &quizService{quizRepository: quizR, animalRepository: animalR}
}

func (qs *quizService) GetAnimalQuiz(ctx context.Context, userID string, animalID string) ([]dto.QuizResponse, error) {
	isCurTarget, err := qs.animalRepository.IsCurrentTarget(ctx, nil, userID, animalID)
	if err != nil {
		return nil, err
	}

	if !isCurTarget {
		return nil, errs.ErrWrongAnimal
	}

	quizzes, err := qs.quizRepository.GetAllQuizOfAnimal(ctx, nil, animalID)
	if err != nil {
		return nil, err
	}

	if len(quizzes) < 3 {
		return nil, errs.ErrQuizCountNotEnough
	}

	var quizzesResp []dto.QuizResponse
	for i := 0; i < 3; i++ {
		randIdx := rand.Intn(len(quizzes))
		quizzesResp = append(quizzesResp, dto.QuizResponse{
			ID:            quizzes[randIdx].ID.String(),
			Question:      quizzes[randIdx].Question,
			CorrectAnswer: quizzes[randIdx].CorrectAnswer,
			WrongAnswer:   quizzes[randIdx].WrongAnswer,
		})
		quizzes[i] = quizzes[len(quizzes)-1]
		quizzes = quizzes[:len(quizzes)-1]
	}
	return quizzesResp, nil
}
