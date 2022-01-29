package exceptions

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"meeting-service/controllers/responses"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(bsoncore.ValidationError)
	fmt.Println(err)
	if ok {
		return ctx.Status(400).JSON(responses.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.Status(500).JSON(responses.WebResponse{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
