package router

import (
	"github.com/zetsux/gin-gorm-clean-starter/api/v1/controller"

	"github.com/gin-gonic/gin"
)

func AnimalTypeRouter(router *gin.Engine, animalTypeC controller.AnimalTypeController) {
	animalTypeRoutes := router.Group("/api/v1/animal-types")
	{
		animalTypeRoutes.GET("/", animalTypeC.GetAllAnimalTypes)
		animalTypeRoutes.GET("/:animal_type_id", animalTypeC.GetAnimalType)
	}
}
