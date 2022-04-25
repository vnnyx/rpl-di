package main

import (
	"rpl-sixmath/config"
	"rpl-sixmath/controller"
	"rpl-sixmath/entity"
	"rpl-sixmath/exception"
	"rpl-sixmath/repository"
	"rpl-sixmath/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configuration := config.New()
	database := config.NewMySQLDatabase(configuration)
	mysqlMaster := config.NewMySQLDatabase(configuration)
	err := mysqlMaster.Debug().AutoMigrate(
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

	userService := service.NewUserService(&userRepository)
	authService := service.NewAuthService(&userRepository)

	userController := controller.NewUserController(&userService)
	authController := controller.NewAuthController(&authService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	userController.Route(app)
	authController.Route(app)

	err = app.Listen(":8000")
	exception.PanicIfNeeded(err)
}
