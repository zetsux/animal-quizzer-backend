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
	GetAllAnimals(ctx context.Context, tx *gorm.DB, filter string) ([]entity.Animal, error)
	GetAllAnimalsByUser(ctx context.Context, tx *gorm.DB, userID string, isOwned bool, filter string) ([]entity.Animal, error)
	GetAnimal(ctx context.Context, tx *gorm.DB, id string) (entity.Animal, error)
	IsCurrentTarget(ctx context.Context, tx *gorm.DB, userID string, animalID string) (bool, error)
}

func NewAnimalRepository(txr *txRepository) *animalRepository {
	return &animalRepository{txr: txr}
}

func (ar *animalRepository) TxRepository() *txRepository {
	return ar.txr
}

func (ar *animalRepository) GetAllAnimals(ctx context.Context, tx *gorm.DB, filter string) ([]entity.Animal, error) {
	var err error
	var animals []entity.Animal
	var stmt *gorm.DB

	if tx == nil {
		stmt = ar.txr.DB().WithContext(ctx).Debug().Preload("AnimalType").Joins("INNER JOIN animal_types ON animals.animal_type_id = animal_types.id")
	} else {
		stmt = tx.WithContext(ctx).Debug().Preload("AnimalType").Joins("INNER JOIN animal_types ON animals.id.animal_type_id = animal_types.id")
	}

	if filter != "" {
		stmt = stmt.Where("animal_types.name = $1", filter)
	}
	err = stmt.Order("animals.name ASC").Find(&animals).Error

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return animals, err
	}
	return animals, nil
}

func (ar *animalRepository) GetAllAnimalsByUser(ctx context.Context, tx *gorm.DB, userID string, isOwned bool, filter string) ([]entity.Animal, error) {
	var err error
	var animals []entity.Animal
	var stmt *gorm.DB

	mark := "<="
	if isOwned {
		mark = ">"
	}

	if tx == nil {
		stmt = ar.txr.DB().WithContext(ctx).Debug().Preload("AnimalType").Joins("INNER JOIN animal_types ON animals.animal_type_id = animal_types.id").Joins("INNER JOIN quests ON animal_types.id = quests.animal_type_id").Where("quests.user_id = $1 AND quests.step "+mark+" animals.step", userID)
	} else {
		stmt = tx.WithContext(ctx).Debug().Preload("AnimalType").Joins("INNER JOIN animal_types ON animals.id.animal_type_id = animal_types.id").Joins("INNER JOIN quests ON animal_types.id = quests.animal_type_id").Where("quests.user_id = $1 AND animals.step "+mark+" quests.step", userID)
	}

	if filter != "" {
		stmt = stmt.Where("animal_types.name = $2", filter)
	}
	err = stmt.Order("animals.name ASC").Find(&animals).Error

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

func (ar *animalRepository) IsCurrentTarget(ctx context.Context, tx *gorm.DB, userID string, animalID string) (bool, error) {
	var err error
	var animal entity.Animal
	if tx == nil {
		tx = ar.txr.DB().WithContext(ctx).Debug().Joins("INNER JOIN animal_types ON animals.animal_type_id = animal_types.id").Joins("INNER JOIN quests ON animal_types.id = quests.animal_type_id").Where("quests.user_id = $1 AND animals.step = quests.step AND animals.id = $2", userID, animalID).Take(&animal)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Joins("INNER JOIN animal_types ON animals.animal_type_id = animal_types.id").Joins("INNER JOIN quests ON animal_types.id = quests.animal_type_id").Where("quests.user_id = $1 AND animals.step = quests.step AND animals.id = $2", userID, animalID).Take(&animal).Error
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
