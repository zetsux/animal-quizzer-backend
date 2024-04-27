package controller

import (
	"net/http"

	"github.com/zetsux/gin-gorm-clean-starter/common/base"
	"github.com/zetsux/gin-gorm-clean-starter/core/helper/messages"
	"github.com/zetsux/gin-gorm-clean-starter/core/service"

	"github.com/gin-gonic/gin"
)

type questController struct {
	questService service.QuestService
}

type QuestController interface {
	GetAllUserQuests(ctx *gin.Context)
	GetUserQuestByAnimalType(ctx *gin.Context)
	AdvanceQuest(ctx *gin.Context)
}

func NewQuestController(questS service.QuestService) QuestController {
	return &questController{
		questService: questS,
	}
}

func (qc *questController) GetAllUserQuests(ctx *gin.Context) {
	userID := ctx.MustGet("ID").(string)
	quests, err := qc.questService.GetAllUserQuests(ctx, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgQuestFetchFailed,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgQuestFetchSuccess,
		http.StatusOK, quests,
	))
}

func (qc *questController) GetUserQuestByAnimalType(ctx *gin.Context) {
	userID := ctx.MustGet("ID").(string)
	animalTypeID := ctx.Param("animal_type_id")

	quest, err := qc.questService.GetUserQuestByAnimalType(ctx, userID, animalTypeID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgQuestFetchFailed,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgQuestFetchSuccess,
		http.StatusOK, quest,
	))
}

func (qc *questController) AdvanceQuest(ctx *gin.Context) {
	userID := ctx.MustGet("ID").(string)
	animalTypeID := ctx.Param("animal_type_id")

	quest, err := qc.questService.AdvanceQuest(ctx, userID, animalTypeID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgQuestUpdateFailed,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgQuestUpdateSuccess,
		http.StatusOK, quest,
	))
}
