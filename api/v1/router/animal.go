package router

import (
	"github.com/zetsux/gin-gorm-clean-starter/api/v1/controller"
	"github.com/zetsux/gin-gorm-clean-starter/common/constant"
	"github.com/zetsux/gin-gorm-clean-starter/common/middleware"
	"github.com/zetsux/gin-gorm-clean-starter/core/service"

	"github.com/gin-gonic/gin"
)

func AnimalRouter(router *gin.Engine, animalC controller.AnimalController, jwtS service.JWTService) {
	animalRoutes := router.Group("/api/v1/animals")
	{
		animalRoutes.GET("/", animalC.GetAllAnimals)
		animalRoutes.GET("/inventory", middleware.Authenticate(jwtS, constant.EnumRoleUser), animalC.GetAnimalInventory)
		animalRoutes.GET("/:animal_id", animalC.GetAnimal)
	}
}
