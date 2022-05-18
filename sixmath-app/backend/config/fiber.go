package config

import (
	"github.com/gofiber/fiber/v2"
	"rpl-sixmath/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
