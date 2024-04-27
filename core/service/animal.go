package service

import (
	"context"
	"reflect"

	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"github.com/zetsux/gin-gorm-clean-starter/core/helper/dto"
	errs "github.com/zetsux/gin-gorm-clean-starter/core/helper/errors"
	"github.com/zetsux/gin-gorm-clean-starter/core/repository"
)

type animalService struct {
	animalRepository repository.AnimalRepository
}

type AnimalService interface {
	GetAllAnimals(ctx context.Context) ([]dto.AnimalResponse, error)
	GetAnimalByID(ctx context.Context, id string) (dto.AnimalResponse, error)
}

func NewAnimalService(animalR repository.AnimalRepository) AnimalService {
	return &animalService{animalRepository: animalR}
}

func (as *animalService) GetAllAnimals(ctx context.Context) (animalsResp []dto.AnimalResponse, err error) {
	animals, err := as.animalRepository.GetAllAnimals(ctx, nil)
	if err != nil {
		return nil, err
	}

	for _, animal := range animals {
		animalsResp = append(animalsResp, dto.AnimalResponse{
			ID:   animal.ID.String(),
			Name: animal.Name,
			Step: animal.Step,
			AnimalType: dto.AnimalTypeResponse{
				ID:   animal.AnimalType.ID.String(),
				Name: animal.AnimalType.Name,
			},
			SilhouetteImage: animal.SilhouetteImage,
			RealImage:       animal.RealImage,
			BadgeImage:      animal.BadgeImage,
		})
	}
	return animalsResp, nil
}

func (as *animalService) GetAnimalByID(ctx context.Context, id string) (dto.AnimalResponse, error) {
	animal, err := as.animalRepository.GetAnimal(ctx, nil, id)
	if err != nil {
		return dto.AnimalResponse{}, err
	}

	if reflect.DeepEqual(animal, entity.Animal{}) {
		return dto.AnimalResponse{}, errs.ErrAnimalNotFound
	}

	return dto.AnimalResponse{
		ID:   animal.ID.String(),
		Name: animal.Name,
		Step: animal.Step,
		Hint: animal.Hint,
		AnimalType: dto.AnimalTypeResponse{
			ID:   animal.AnimalType.ID.String(),
			Name: animal.AnimalType.Name,
		},
		SilhouetteImage: animal.SilhouetteImage,
		RealImage:       animal.RealImage,
		BadgeImage:      animal.BadgeImage,
	}, nil
}
