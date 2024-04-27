package repository

import (
	"context"
	"errors"

	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"gorm.io/gorm"
)

type quizRepository struct {
	txr *txRepository
}

type QuizRepository interface {
	// tx
	TxRepository() *txRepository

	// functional
	GetQuiz(ctx context.Context, tx *gorm.DB, id string) (entity.Quiz, error)
	GetAllQuizOfAnimal(ctx context.Context, tx *gorm.DB, animal_id string) ([]entity.Quiz, error)
}

func NewQuizRepository(txr *txRepository) *quizRepository {
	return &quizRepository{txr: txr}
}

func (ar *quizRepository) TxRepository() *txRepository {
	return ar.txr
}

func (ar *quizRepository) GetQuiz(ctx context.Context, tx *gorm.DB, id string) (entity.Quiz, error) {
	var err error
	var quiz entity.Quiz
	if tx == nil {
		tx = ar.txr.DB().WithContext(ctx).Debug().Where("id = $1", id).Take(&quiz)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Where("id = $1", id).Take(&quiz).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return quiz, err
	}
	return quiz, nil
}

func (ar *quizRepository) GetAllQuizOfAnimal(ctx context.Context, tx *gorm.DB, animal_id string) ([]entity.Quiz, error) {
	var err error
	var quizzes []entity.Quiz
	if tx == nil {
		tx = ar.txr.DB().WithContext(ctx).Debug().Where("animal_id = $1", animal_id).Find(&quizzes)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Where("animal_id = $1", animal_id).Find(&quizzes).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return quizzes, err
	}
	return quizzes, nil
}
