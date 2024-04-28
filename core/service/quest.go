package service

import (
	"context"

	"github.com/zetsux/gin-gorm-clean-starter/core/helper/dto"
	"github.com/zetsux/gin-gorm-clean-starter/core/repository"
)

type questService struct {
	questRepository  repository.QuestRepository
	animalRepository repository.AnimalRepository
}

type QuestService interface {
	GetAllUserQuests(ctx context.Context, userID string) ([]dto.QuestResponse, error)
	GetUserQuestByAnimalType(ctx context.Context, userID string, animalTypeID string) (dto.QuestResponse, error)
	AdvanceQuest(ctx context.Context, userID string, animalID string) (dto.QuestResponse, error)
	GetQuestLeaderboard(ctx context.Context, isDaily bool) ([]dto.QuestLeaderboard, error)
}

func NewQuestService(questR repository.QuestRepository, animalR repository.AnimalRepository) QuestService {
	return &questService{questRepository: questR, animalRepository: animalR}
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

	animal, err := qs.animalRepository.GetCurrentTarget(ctx, nil, animalTypeID, quest.Step)
	if err != nil {
		return dto.QuestResponse{}, err
	}

	return dto.QuestResponse{
		ID:              quest.ID.String(),
		Step:            quest.Step,
		Hint:            animal.Hint,
		SilhouetteImage: animal.SilhouetteImage,
		AnimalType: dto.AnimalTypeResponse{
			ID:   quest.AnimalType.ID.String(),
			Name: quest.AnimalType.Name,
		},
	}, nil
}

func (qs *questService) AdvanceQuest(ctx context.Context, userID string, animalID string) (dto.QuestResponse, error) {
	quest, err := qs.questRepository.GetUserQuestByAnimal(ctx, nil, animalID, userID)
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

func (qs *questService) GetQuestLeaderboard(ctx context.Context, isDaily bool) ([]dto.QuestLeaderboard, error) {
	leaderboard, err := qs.questRepository.GetQuestLeaderboard(ctx, nil, isDaily)
	if err != nil {
		return nil, err
	}
	return leaderboard, nil
}
