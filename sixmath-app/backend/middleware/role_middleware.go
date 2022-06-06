package middleware

import (
	"net/http"
	"rpl-sixmath/model"

	"github.com/gofiber/fiber/v2"
)

func IsAdmin() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Locals("currentRole") != "admin" {
			response := model.Response{
				Code:   403,
				Status: "Unauthorized",
				Error: map[string]interface{}{
					"general": "UNAUTHORIZED",
				},
			}
			return c.Status(http.StatusForbidden).JSON(response)
		}

		return c.Next()
	}
}

func IsParent() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Locals("currentRole") != "parent" {
			response := model.Response{
				Code:   403,
				Status: "Unauthorized",
				Error: map[string]interface{}{
					"general": "UNAUTHORIZED",
				},
			}
			return c.Status(http.StatusForbidden).JSON(response)
		}

		return c.Next()
	}
}

func IsStudent() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Locals("currentRole") != "student" {
			response := model.Response{
				Code:   403,
				Status: "Unauthorized",
				Error: map[string]interface{}{
					"general": "UNAUTHORIZED",
				},
			}
			return c.Status(http.StatusForbidden).JSON(response)
		}

		return c.Next()
	}
}

func IsTeacher() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Locals("currentRole") != "teacher" {
			response := model.Response{
				Code:   403,
				Status: "Unauthorized",
				Error: map[string]interface{}{
					"general": "UNAUTHORIZED",
				},
			}
			return c.Status(http.StatusForbidden).JSON(response)
		}

		return c.Next()
	}
}
