package main

import (
	"fmt"
	"os"

	"github.com/zetsux/gin-gorm-clean-starter/api/v1/controller"
	"github.com/zetsux/gin-gorm-clean-starter/api/v1/router"
	"github.com/zetsux/gin-gorm-clean-starter/common/middleware"
	"github.com/zetsux/gin-gorm-clean-starter/config"
	"github.com/zetsux/gin-gorm-clean-starter/core/repository"
	"github.com/zetsux/gin-gorm-clean-starter/core/service"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		db   = config.DBSetup()
		jwtS = service.NewJWTService()
		txR  = repository.NewTxRepository(db)

		fileC = controller.NewFileController()

		animalTypeR = repository.NewAnimalTypeRepository(txR)
		animalTypeS = service.NewAnimalTypeService(animalTypeR)
		animalTypeC = controller.NewAnimalTypeController(animalTypeS)

		animalR = repository.NewAnimalRepository(txR)
		animalS = service.NewAnimalService(animalR)
		animalC = controller.NewAnimalController(animalS)

		questR = repository.NewQuestRepository(txR)
		questS = service.NewQuestService(questR)
		questC = controller.NewQuestController(questS)

		quizR = repository.NewQuizRepository(txR)
		quizS = service.NewQuizService(quizR, animalR)
		quizC = controller.NewQuizController(quizS)

		userR = repository.NewUserRepository(txR)
		userS = service.NewUserService(userR, animalTypeR, questR)
		userC = controller.NewUserController(userS, jwtS)
	)

	defer config.DBClose(db)

	// Setting Up Server
	server := gin.Default()
	server.Use(
		middleware.CORSMiddleware(),
	)

	// Setting Up Routes
	router.UserRouter(server, userC, jwtS)
	router.FileRouter(server, fileC)
	router.AnimalTypeRouter(server, animalTypeC)
	router.AnimalRouter(server, animalC, jwtS)
	router.QuestRouter(server, questC, jwtS)
	router.QuizRouter(server, quizC, jwtS)

	// Running in localhost:8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := server.Run(":" + port)
	if err != nil {
		fmt.Println("Server failed to start: ", err)
		return
	}
}
