package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zetsux/gin-gorm-clean-starter/common/base"
	"github.com/zetsux/gin-gorm-clean-starter/core/helper/messages"
	"github.com/zetsux/gin-gorm-clean-starter/core/service"
)

type quizController struct {
	quizService service.QuizService
}

type QuizController interface {
	GetAnimalQuiz(ctx *gin.Context)
}

func NewQuizController(quizS service.QuizService) QuizController {
	return &quizController{
		quizService: quizS,
	}
}

func (qc *quizController) GetAnimalQuiz(ctx *gin.Context) {
	userID := ctx.MustGet("ID").(string)
	animalID := ctx.Param("animal_id")

	quizs, err := qc.quizService.GetAnimalQuiz(ctx, userID, animalID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgQuizFetchFailed,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgQuizFetchSuccess,
		http.StatusOK, quizs,
	))
}
