package service

import (
	"context"
	"reflect"

	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"github.com/zetsux/gin-gorm-clean-starter/core/helper/dto"
	errs "github.com/zetsux/gin-gorm-clean-starter/core/helper/errors"
	"github.com/zetsux/gin-gorm-clean-starter/core/repository"
)

type animalTypeService struct {
	animalTypeRepository repository.AnimalTypeRepository
}

type AnimalTypeService interface {
	GetAllAnimalTypes(ctx context.Context) ([]dto.AnimalTypeResponse, error)
	GetAnimalTypeByID(ctx context.Context, id string) (dto.AnimalTypeResponse, error)
}

func NewAnimalTypeService(animalTypeR repository.AnimalTypeRepository) AnimalTypeService {
	return &animalTypeService{animalTypeRepository: animalTypeR}
}

func (ats *animalTypeService) GetAllAnimalTypes(ctx context.Context) (animalTypesResp []dto.AnimalTypeResponse, err error) {
	animalTypes, err := ats.animalTypeRepository.GetAllAnimalTypes(ctx, nil)
	if err != nil {
		return nil, err
	}

	for _, animalType := range animalTypes {
		animalTypesResp = append(animalTypesResp, dto.AnimalTypeResponse{
			ID:   animalType.ID.String(),
			Name: animalType.Name,
		})
	}
	return animalTypesResp, nil
}

func (ats *animalTypeService) GetAnimalTypeByID(ctx context.Context, id string) (dto.AnimalTypeResponse, error) {
	animalType, err := ats.animalTypeRepository.GetAnimalType(ctx, nil, id)
	if err != nil {
		return dto.AnimalTypeResponse{}, err
	}

	if reflect.DeepEqual(animalType, entity.AnimalType{}) {
		return dto.AnimalTypeResponse{}, errs.ErrAnimalTypeNotFound
	}

	return dto.AnimalTypeResponse{
		ID:   animalType.ID.String(),
		Name: animalType.Name,
	}, nil
}
