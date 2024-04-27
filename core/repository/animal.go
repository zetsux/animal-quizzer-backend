package repository

import (
	"context"
	"errors"

	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"gorm.io/gorm"
)

type animalRepository struct {
	txr *txRepository
}

type AnimalRepository interface {
	// tx
	TxRepository() *txRepository

	// functional
	GetAllAnimals(ctx context.Context, tx *gorm.DB) ([]entity.Animal, error)
	GetAnimal(ctx context.Context, tx *gorm.DB, id string) (entity.Animal, error)
}

func NewAnimalRepository(txr *txRepository) *animalRepository {
	return &animalRepository{txr: txr}
}

func (ar *animalRepository) TxRepository() *txRepository {
	return ar.txr
}

func (ar *animalRepository) GetAllAnimals(ctx context.Context, tx *gorm.DB) ([]entity.Animal, error) {
	var err error
	var animals []entity.Animal
	if tx == nil {
		tx = ar.txr.DB().WithContext(ctx).Debug().Preload("AnimalType").Find(&animals)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Preload("AnimalType").Find(&animals).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return animals, err
	}
	return animals, nil
}

func (ar *animalRepository) GetAnimal(ctx context.Context, tx *gorm.DB, id string) (entity.Animal, error) {
	var err error
	var animal entity.Animal
	if tx == nil {
		tx = ar.txr.DB().WithContext(ctx).Debug().Preload("AnimalType").Where("id = $1", id).Take(&animal)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Preload("AnimalType").Where("id = $1", id).Take(&animal).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return animal, err
	}
	return animal, nil
}
