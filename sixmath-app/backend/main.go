package main

import (
	"rpl-sixmath/config"
	"rpl-sixmath/controller"
	"rpl-sixmath/entity"
	"rpl-sixmath/exception"
	"rpl-sixmath/migration"
	"rpl-sixmath/repository"
	"rpl-sixmath/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configuration := config.New(".env")
	databases := config.NewMySQLDatabase(configuration)
	migration.MyMigration(databases,
		entity.User{},
		entity.Message{},
		entity.Transaction{},
		entity.Playlist{},
		entity.Video{},
		entity.Exam{},
		entity.Question{},
		entity.Answer{},
	)

	userRepository := repository.NewUserRepository(databases)
	playlistRepository := repository.NewPlaylistRepository(databases)
	videoRepository := repository.NewVideoRepository(databases)
	examRepository := repository.NewExamRepository(databases)

	userService := service.NewUserService(&userRepository)
	authService := service.NewAuthService(&userRepository)
	playlistService := service.NewPlaylistService(&playlistRepository)
	videoService := service.NewVideoService(&videoRepository)
	examService := service.NewExamService(examRepository, videoRepository)

	userController := controller.NewUserController(&userService)
	authController := controller.NewAuthController(&authService)
	playlistController := controller.NewPlaylistController(&playlistService)
	videoController := controller.NewVideoController(&videoService)
	examController := controller.NewExamController(examService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(cors.New())
	app.Use(recover.New())

	userController.Route(app)
	authController.Route(app)
	playlistController.Route(app)
	videoController.Route(app)
	examController.Route(app)

	err := app.Listen(":8000")
	exception.PanicIfNeeded(err)
}
