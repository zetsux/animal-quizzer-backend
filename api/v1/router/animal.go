package router

import (
	"github.com/zetsux/gin-gorm-clean-starter/api/v1/controller"

	"github.com/gin-gonic/gin"
)

func AnimalRouter(router *gin.Engine, animalC controller.AnimalController) {
	animalRoutes := router.Group("/api/v1/animals")
	{
		animalRoutes.GET("/", animalC.GetAllAnimals)
		animalRoutes.GET("/:animal_id", animalC.GetAnimal)
	}
}
