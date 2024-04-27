package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zetsux/gin-gorm-clean-starter/common/base"
	"github.com/zetsux/gin-gorm-clean-starter/core/helper/messages"
	"github.com/zetsux/gin-gorm-clean-starter/core/service"
)

type animalController struct {
	animalService service.AnimalService
}

type AnimalController interface {
	GetAllAnimals(ctx *gin.Context)
	GetAnimalInventory(ctx *gin.Context)
	GetAnimal(ctx *gin.Context)
}

func NewAnimalController(animalS service.AnimalService) AnimalController {
	return &animalController{
		animalService: animalS,
	}
}

func (atc *animalController) GetAllAnimals(ctx *gin.Context) {
	animals, err := atc.animalService.GetAllAnimals(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgAnimalFetchFailed,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgAnimalFetchSuccess,
		http.StatusOK, animals,
	))
}

func (atc *animalController) GetAnimalInventory(ctx *gin.Context) {
	userID := ctx.MustGet("ID").(string)
	filter := ctx.Query("filter")

	animals, err := atc.animalService.GetAnimalInventory(ctx, userID, filter)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgAnimalFetchFailed,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgAnimalFetchSuccess,
		http.StatusOK, animals,
	))
}

func (atc *animalController) GetAnimal(ctx *gin.Context) {
	id := ctx.Param("animal_id")

	animal, err := atc.animalService.GetAnimalByID(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgAnimalFetchFailed,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgAnimalFetchSuccess,
		http.StatusOK, animal,
	))
}
