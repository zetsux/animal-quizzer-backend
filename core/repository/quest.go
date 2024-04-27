package repository

import (
	"context"
	"errors"
	"time"

	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"github.com/zetsux/gin-gorm-clean-starter/core/helper/dto"
	"gorm.io/gorm"
)

type questRepository struct {
	txr *txRepository
}

type QuestRepository interface {
	// tx
	TxRepository() *txRepository

	// functional
	GetAllUserQuests(ctx context.Context, tx *gorm.DB, userID string) ([]entity.Quest, error)
	GetUserQuestByAnimalType(ctx context.Context, tx *gorm.DB, animalTypeID string, userID string) (entity.Quest, error)
	CreateNewQuest(ctx context.Context, tx *gorm.DB, quest entity.Quest) (entity.Quest, error)
	UpdateQuest(ctx context.Context, tx *gorm.DB, quest entity.Quest) (entity.Quest, error)
	GetQuestLeaderboard(ctx context.Context, tx *gorm.DB, isDaily bool) ([]dto.QuestLeaderboard, error)
}

func NewQuestRepository(txr *txRepository) *questRepository {
	return &questRepository{txr: txr}
}

func (qr *questRepository) TxRepository() *txRepository {
	return qr.txr
}

func (qr *questRepository) GetAllUserQuests(ctx context.Context, tx *gorm.DB, userID string) ([]entity.Quest, error) {
	var err error
	var quests []entity.Quest
	if tx == nil {
		tx = qr.txr.DB().WithContext(ctx).Debug().Preload("AnimalType").Where("user_id = $1", userID).Find(&quests)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Preload("AnimalType").Where("user_id = $1", userID).Find(&quests).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return quests, err
	}
	return quests, nil
}

func (qr *questRepository) GetUserQuestByAnimalType(ctx context.Context, tx *gorm.DB, animalTypeID string, userID string) (entity.Quest, error) {
	var err error
	var quest entity.Quest
	if tx == nil {
		tx = qr.txr.DB().WithContext(ctx).Debug().Preload("AnimalType").Where("animal_type_id = $1 AND user_id = $2", animalTypeID, userID).Find(&quest)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Preload("AnimalType").Where("animal_type_id = $1 AND user_id = $2", animalTypeID, userID).Find(&quest).Error
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return quest, err
	}
	return quest, nil
}

func (ur *questRepository) CreateNewQuest(ctx context.Context, tx *gorm.DB, quest entity.Quest) (entity.Quest, error) {
	var err error
	if tx == nil {
		tx = ur.txr.DB().WithContext(ctx).Debug().Create(&quest)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Create(&quest).Error
	}

	if err != nil {
		return entity.Quest{}, err
	}
	return quest, nil
}

func (qr *questRepository) UpdateQuest(ctx context.Context, tx *gorm.DB, quest entity.Quest) (entity.Quest, error) {
	if tx == nil {
		tx = qr.txr.DB().WithContext(ctx).Debug().Updates(&quest)
		if err := tx.Error; err != nil {
			return entity.Quest{}, err
		}
	} else {
		if err := qr.txr.DB().Updates(&quest).Error; err != nil {
			return entity.Quest{}, err
		}
	}

	return quest, nil
}

func (qr *questRepository) GetQuestLeaderboard(ctx context.Context, tx *gorm.DB, isDaily bool) ([]dto.QuestLeaderboard, error) {
	var err error
	var quests []dto.QuestLeaderboard
	var stmt *gorm.DB

	if tx == nil {
		stmt = qr.txr.DB().WithContext(ctx).Debug().Model(&entity.User{}).Select("users.username as username", "(users.last_attempt - users.created_at) as time", "sum(quests.step) as point").Joins("inner join quests on quests.user_id = users.id")
	} else {
		stmt = tx.WithContext(ctx).Debug().Model(&entity.User{}).Select("users.username as username", "(users.last_attempt - users.created_at) as time", "sum(quests.step) as point").Joins("inner join quests on quests.user_id = users.id")
	}

	if isDaily {
		today := time.Now()
		beginningOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
		stmt = stmt.Where("users.created_at > $1", beginningOfDay)
	}

	err = stmt.Group("users.id").Order("point desc").Order("time").Scan(&quests).Error
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return quests, err
	}
	return quests, nil
}
