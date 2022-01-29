package controllers

import (
	"meeting-service/controllers/mappers"
	"meeting-service/controllers/requests"
	"meeting-service/controllers/responses"
	"meeting-service/exceptions"
	"meeting-service/usecases/in"

	"github.com/gofiber/fiber/v2"
)

type MeetingController struct {
	create  in.CreateMeetingInput
	details in.RetrieveMeetingDetailsInput
}

func NewMeetingController(
	create *in.CreateMeetingInput,
	details *in.RetrieveMeetingDetailsInput,
) MeetingController {
	return MeetingController{
		*create,
		*details,
	}
}

func (controller *MeetingController) Route(app *fiber.App) {
	app.Post("/meetings", controller.Create)
	app.Get("/meetings/:meetingId", controller.Detail)
}

func (controller *MeetingController) Create(c *fiber.Ctx) error {
	var request requests.CreateMeetingRequest
	err := c.BodyParser(&request)
	var meetingEntity = mappers.CreatePostRequestToPostMapper(request)

	exceptions.PanicIfNeeded(err)

	controller.create.Create(*meetingEntity)
	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meetingEntity.Id,
	})
}

func (controller *MeetingController) Detail(c *fiber.Ctx) error {
	result := controller.details.Get(c.Params("meetingId"))

	return c.JSON(responses.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	})
}
