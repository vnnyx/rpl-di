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
		entity.Video{},
		entity.Exam{},
		entity.Question{},
		entity.Answer{},
		entity.PasswordReset{},
	)

	userRepository := repository.NewUserRepository(databases)
	videoRepository := repository.NewVideoRepository(databases)
	examRepository := repository.NewExamRepository(databases)
	passwordResetRepository := repository.NewPasswordResetRepository(databases)

	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(&userRepository)
	videoService := service.NewVideoService(videoRepository)
	examService := service.NewExamService(examRepository, videoRepository)
	passwordResetService := service.NewPasswordResetService(passwordResetRepository, userRepository)

	userController := controller.NewUserController(&userService)
	authController := controller.NewAuthController(&authService)
	videoController := controller.NewVideoController(&videoService)
	examController := controller.NewExamController(examService)
	passwordResetController := controller.NewPasswordResetController(&passwordResetService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(cors.New())
	app.Use(recover.New())

	userController.Route(app)
	authController.Route(app)
	videoController.Route(app)
	examController.Route(app)
	passwordResetController.Route(app)

	err := app.Listen(":8000")
	exception.PanicIfNeeded(err)
}
