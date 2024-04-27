package router

import (
	"github.com/zetsux/gin-gorm-clean-starter/api/v1/controller"
	"github.com/zetsux/gin-gorm-clean-starter/common/constant"
	"github.com/zetsux/gin-gorm-clean-starter/common/middleware"
	"github.com/zetsux/gin-gorm-clean-starter/core/service"

	"github.com/gin-gonic/gin"
)

func QuizRouter(router *gin.Engine, quizC controller.QuizController, jwtS service.JWTService) {
	quizRoutes := router.Group("/api/v1/quizzes")
	{
		quizRoutes.GET("/:animal_id", middleware.Authenticate(jwtS, constant.EnumRoleUser), quizC.GetAnimalQuiz)
		quizRoutes.PATCH("/", middleware.Authenticate(jwtS, constant.EnumRoleUser), quizC.SetQuizCooldown)
	}
}
