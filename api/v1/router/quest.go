package router

import (
	"github.com/zetsux/gin-gorm-clean-starter/api/v1/controller"
	"github.com/zetsux/gin-gorm-clean-starter/common/constant"
	"github.com/zetsux/gin-gorm-clean-starter/common/middleware"
	"github.com/zetsux/gin-gorm-clean-starter/core/service"

	"github.com/gin-gonic/gin"
)

func QuestRouter(router *gin.Engine, questC controller.QuestController, jwtS service.JWTService) {
	questRoutes := router.Group("/api/v1/quests")
	{
		questRoutes.GET("/", middleware.Authenticate(jwtS, constant.EnumRoleUser), questC.GetAllUserQuests)
		questRoutes.GET("/:animal_type_id", middleware.Authenticate(jwtS, constant.EnumRoleUser), questC.GetUserQuestByAnimalType)
		questRoutes.PATCH("/:animal_type_id", middleware.Authenticate(jwtS, constant.EnumRoleUser), questC.AdvanceQuest)
		questRoutes.GET("/leaderboard", questC.GetQuestLeaderboard)
		questRoutes.GET("/daily-leaderboard", questC.GetDailyQuestLeaderboard)
	}
}
