package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zetsux/gin-gorm-clean-starter/common/base"
	errs "github.com/zetsux/gin-gorm-clean-starter/core/helper/errors"
	"github.com/zetsux/gin-gorm-clean-starter/core/helper/messages"
	"github.com/zetsux/gin-gorm-clean-starter/core/service"
)

type quizController struct {
	quizService service.QuizService
}

type QuizController interface {
	GetAnimalQuiz(ctx *gin.Context)
	SetQuizCooldown(ctx *gin.Context)
}

func NewQuizController(quizS service.QuizService) QuizController {
	return &quizController{
		quizService: quizS,
	}
}

func (qc *quizController) GetAnimalQuiz(ctx *gin.Context) {
	userID := ctx.MustGet("ID").(string)
	animalID := ctx.Param("animal_id")

	quizzes, err := qc.quizService.GetAnimalQuiz(ctx, userID, animalID)
	if err != nil {
		if errors.Is(err, errs.ErrQuizInCooldown) {
			ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
				messages.MsgQuizInCooldown,
				http.StatusOK, quizzes[0],
			))
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
				messages.MsgQuizFetchFailed,
				err.Error(), http.StatusBadRequest,
			))
		}
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgQuizFetchSuccess,
		http.StatusOK, quizzes,
	))
}

func (qc *quizController) SetQuizCooldown(ctx *gin.Context) {
	userID := ctx.MustGet("ID").(string)

	err := qc.quizService.SetCooldown(ctx, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgQuizSetCooldownFailed,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgQuizSetCooldownSuccess,
		http.StatusOK, nil,
	))
}
