package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zetsux/gin-gorm-clean-starter/common/base"
	"github.com/zetsux/gin-gorm-clean-starter/core/helper/messages"
	"github.com/zetsux/gin-gorm-clean-starter/core/service"
)

type animalTypeController struct {
	animalTypeService service.AnimalTypeService
}

type AnimalTypeController interface {
	GetAllAnimalTypes(ctx *gin.Context)
	GetAnimalType(ctx *gin.Context)
}

func NewAnimalTypeController(animalTypeS service.AnimalTypeService) AnimalTypeController {
	return &animalTypeController{
		animalTypeService: animalTypeS,
	}
}

func (atc *animalTypeController) GetAllAnimalTypes(ctx *gin.Context) {
	animalTypes, err := atc.animalTypeService.GetAllAnimalTypes(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgAnimalTypeFetchFailed,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgAnimalTypeFetchSuccess,
		http.StatusOK, animalTypes,
	))
}

func (atc *animalTypeController) GetAnimalType(ctx *gin.Context) {
	id := ctx.Param("animal_type_id")

	animalType, err := atc.animalTypeService.GetAnimalTypeByID(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse(
			messages.MsgAnimalTypeFetchFailed,
			err.Error(), http.StatusBadRequest,
		))
		return
	}

	ctx.JSON(http.StatusOK, base.CreateSuccessResponse(
		messages.MsgAnimalTypeFetchSuccess,
		http.StatusOK, animalType,
	))
}
