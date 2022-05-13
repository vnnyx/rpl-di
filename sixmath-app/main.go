package main

import (
	"rpl-sixmath/config"
	"rpl-sixmath/controller"
	"rpl-sixmath/entity"
	"rpl-sixmath/exception"
	"rpl-sixmath/repository"
	"rpl-sixmath/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configuration := config.New()
	database := config.NewMySQLDatabase(configuration)
	err := database.Debug().AutoMigrate(
		entity.UserEntity{},
		entity.MessageEntity{},
		entity.TransactionEntity{},
		entity.PlaylistEntity{},
		entity.VideoEntity{},
		entity.ExamEntity{},
		entity.QuestionEntity{},
		entity.AnswerEntity{},
	)
	exception.PanicIfNeeded(err)

	userRepository := repository.NewUserRepository(database)
	playlistRepository := repository.NewPlaylistRepository(database)
	videoRepository := repository.NewVideoRepository(database)

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
