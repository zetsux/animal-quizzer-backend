package service

import (
	"context"
	"reflect"

	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"github.com/zetsux/gin-gorm-clean-starter/core/helper/dto"
	"github.com/zetsux/gin-gorm-clean-starter/core/repository"
)

type questService struct {
	questRepository repository.QuestRepository
}

type QuestService interface {
	GetAllUserQuests(ctx context.Context, userID string) ([]dto.QuestResponse, error)
	GetUserQuestByAnimalType(ctx context.Context, userID string, animalTypeID string) (dto.QuestResponse, error)
	AdvanceQuest(ctx context.Context, userID string, animalTypeID string) (dto.QuestResponse, error)
}

func NewQuestService(questR repository.QuestRepository) QuestService {
	return &questService{questRepository: questR}
}

func (qs *questService) GetAllUserQuests(ctx context.Context, userID string) (questsResp []dto.QuestResponse, err error) {
	quests, err := qs.questRepository.GetAllUserQuests(ctx, nil, userID)
	if err != nil {
		return nil, err
	}

	for _, quest := range quests {
		questsResp = append(questsResp, dto.QuestResponse{
			ID:   quest.ID.String(),
			Step: quest.Step,
			AnimalType: dto.AnimalTypeResponse{
				ID:   quest.AnimalType.ID.String(),
				Name: quest.AnimalType.Name,
			},
		})
	}
	return questsResp, nil
}

func (qs *questService) GetUserQuestByAnimalType(ctx context.Context, userID string, animalTypeID string) (dto.QuestResponse, error) {
	quest, err := qs.questRepository.GetUserQuestByAnimalType(ctx, nil, animalTypeID, userID)
	if err != nil {
		return dto.QuestResponse{}, err
	}

	if reflect.DeepEqual(quest, entity.Quest{}) {
		quest = entity.Quest{
			Step:         1,
			UserID:       userID,
			AnimalTypeID: animalTypeID,
		}

		// create new quest
		quest, err = qs.questRepository.CreateNewQuest(ctx, nil, quest)
		if err != nil {
			return dto.QuestResponse{}, err
		}

		quest, err = qs.questRepository.GetUserQuestByAnimalType(ctx, nil, animalTypeID, userID)
		if err != nil {
			return dto.QuestResponse{}, err
		}
	}

	return dto.QuestResponse{
		ID:   quest.ID.String(),
		Step: quest.Step,
		AnimalType: dto.AnimalTypeResponse{
			ID:   quest.AnimalType.ID.String(),
			Name: quest.AnimalType.Name,
		},
	}, nil
}

func (qs *questService) AdvanceQuest(ctx context.Context, userID string, animalTypeID string) (dto.QuestResponse, error) {
	quest, err := qs.questRepository.GetUserQuestByAnimalType(ctx, nil, animalTypeID, userID)
	if err != nil {
		return dto.QuestResponse{}, err
	}
	quest.Step += 1

	quest, err = qs.questRepository.UpdateQuest(ctx, nil, quest)
	if err != nil {
		return dto.QuestResponse{}, err
	}

	return dto.QuestResponse{
		ID:   quest.ID.String(),
		Step: quest.Step,
		AnimalType: dto.AnimalTypeResponse{
			ID:   quest.AnimalType.ID.String(),
			Name: quest.AnimalType.Name,
		},
	}, nil
}
