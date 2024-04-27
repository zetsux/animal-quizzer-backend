package repository

import (
	"context"
	"errors"

	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"gorm.io/gorm"
)

type animalTypeRepository struct {
	txr *txRepository
}

type AnimalTypeRepository interface {
	// tx
	TxRepository() *txRepository

	// functional
	GetAllAnimalTypes(ctx context.Context, tx *gorm.DB) ([]entity.AnimalType, error)
	GetAnimalType(ctx context.Context, tx *gorm.DB, id string) (entity.AnimalType, error)
}

func NewAnimalTypeRepository(txr *txRepository) *animalTypeRepository {
	return &animalTypeRepository{txr: txr}
}

func (atr *animalTypeRepository) TxRepository() *txRepository {
	return atr.txr
}

func (atr *animalTypeRepository) GetAllAnimalTypes(ctx context.Context, tx *gorm.DB) ([]entity.AnimalType, error) {
	var err error
	var animalTypes []entity.AnimalType
	if tx == nil {
		tx = atr.txr.DB().WithContext(ctx).Debug().Find(&animalTypes)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Find(&animalTypes).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return animalTypes, err
	}
	return animalTypes, nil
}

func (atr *animalTypeRepository) GetAnimalType(ctx context.Context, tx *gorm.DB, id string) (entity.AnimalType, error) {
	var err error
	var animalType entity.AnimalType
	if tx == nil {
		tx = atr.txr.DB().WithContext(ctx).Debug().Where("id = $1", id).Take(&animalType)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Where("id = $1", id).Take(&animalType).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return animalType, err
	}
	return animalType, nil
}
