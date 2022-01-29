package configurations

import (
	"github.com/gofiber/fiber/v2"
	"meeting-service/exceptions"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exceptions.ErrorHandler,
	}
}
