package controller

import (
	"rpl-sixmath/config"
	"rpl-sixmath/repository"
	"rpl-sixmath/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func createTestApp() *fiber.App {
	//migration.MyMigration(database,
	//	entity.User{},
	//	entity.Message{},
	//	entity.Transaction{},
	//	entity.Playlist{},
	//	entity.Video{},
	//	entity.Exam{},
	//	entity.Question{},
	//	entity.Answer{},
	//)
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	userController.Route(app)
	return app
}

var configuration = config.New("../.env.test")

var database = config.NewMySQLDatabase(configuration)
var userRepository = repository.NewUserRepository(database)
var userService = service.NewUserService(userRepository)
var userController = NewUserController(&userService)
var app = createTestApp()
