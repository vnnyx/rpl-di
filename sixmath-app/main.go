package main

import (
	"rpl-sixmath/config"
	"rpl-sixmath/controller"
	"rpl-sixmath/exception"
	"rpl-sixmath/model/database"
	"rpl-sixmath/repository"
	"rpl-sixmath/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configuration := config.New()
	databases := config.NewMySQLDatabase(configuration)
	err := databases.Debug().AutoMigrate(
		database.UserEntity{},
		database.MessageEntity{},
		database.TransactionEntity{},
		database.PlaylistEntity{},
		database.VideoEntity{},
		database.ExamEntity{},
		database.QuestionEntity{},
		database.AnswerEntity{},
	)
	exception.PanicIfNeeded(err)

	userRepository := repository.NewUserRepository(databases)
	playlistRepository := repository.NewPlaylistRepository(databases)
	videoRepository := repository.NewVideoRepository(databases)

	userService := service.NewUserService(&userRepository)
	authService := service.NewAuthService(&userRepository)
	playlistService := service.NewPlaylistService(&playlistRepository)
	videoService := service.NewVideoService(&videoRepository)

	userController := controller.NewUserController(&userService)
	authController := controller.NewAuthController(&authService)
	playlistController := controller.NewPlaylistController(&playlistService)
	videoController := controller.NewVideoController(&videoService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(cors.New())
	app.Use(recover.New())

	userController.Route(app)
	authController.Route(app)
	playlistController.Route(app)
	videoController.Route(app)

	err = app.Listen(":8000")
	exception.PanicIfNeeded(err)
}
