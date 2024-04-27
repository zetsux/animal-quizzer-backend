package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"github.com/zetsux/gin-gorm-clean-starter/core/helper/dto"
	errs "github.com/zetsux/gin-gorm-clean-starter/core/helper/errors"
	"github.com/zetsux/gin-gorm-clean-starter/core/repository"
)

type quizService struct {
	quizRepository   repository.QuizRepository
	animalRepository repository.AnimalRepository
	userRepository   repository.UserRepository
}

type QuizService interface {
	GetAnimalQuiz(ctx context.Context, userID string, animalID string) ([]dto.QuizResponse, error)
	SetCooldown(ctx context.Context, userID string) error
}

func NewQuizService(quizR repository.QuizRepository, animalR repository.AnimalRepository, userR repository.UserRepository) QuizService {
	return &quizService{quizRepository: quizR, animalRepository: animalR, userRepository: userR}
}

func (qs *quizService) GetAnimalQuiz(ctx context.Context, userID string, animalID string) ([]dto.QuizResponse, error) {
	isCurTarget, err := qs.animalRepository.IsCurrentTarget(ctx, nil, userID, animalID)
	if err != nil {
		return nil, err
	}

	if !isCurTarget {
		return nil, errs.ErrWrongAnimal
	}

	user, err := qs.userRepository.GetUserByPrimaryKey(ctx, nil, "id", userID)
	if err != nil {
		return nil, err
	}

	if time.Since(user.LastAttempt) < (time.Minute * 3) {
		return []dto.QuizResponse{
			{
				Cooldown: int(((3 * time.Minute) - time.Since(user.LastAttempt)).Seconds()),
			},
		}, errs.ErrQuizInCooldown
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

func (qs *quizService) SetCooldown(ctx context.Context, userID string) error {
	_, err := qs.userRepository.UpdateUser(ctx, nil, entity.User{
		ID:          uuid.MustParse(userID),
		LastAttempt: time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}
