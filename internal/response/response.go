package response

import (
	"github.com/gofiber/fiber/v3"
)

type APIError struct {
	Type        string `json:"type"`
	Msg         string `json:"msg"`
	ErrorNumber int    `json:"error_number"`
}

func SendError(c fiber.Ctx, code int, payload APIError) error {
	// Toto vraci null - co s tim?
	return c.Status(fiber.StatusBadRequest).JSON(payload)
}
