package main

import (
	"fmt"
	"meeting-service/configurations"
	"meeting-service/controllers"
	"meeting-service/exceptions"
	"meeting-service/repositories/mongo"
	"meeting-service/usecases"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	fmt.Println("Setting up datasource")
	database := configurations.NewMongoDatabase()

	fmt.Println("Setting up repositories")
	meetingRepository := mongo.NewMeetingRepository(database)

	fmt.Println("Setting up use cases")
	retrieveMeeting := usecases.NewRetrieveMeetingDetailsInteractor(
		&meetingRepository,
	)
	createMeeting := usecases.NewCreateMeetingInput(
		&meetingRepository,
	)

	fmt.Println("Setting up controllers")
	meetingController := controllers.NewMeetingController(
		&createMeeting,
		&retrieveMeeting,
	)

	app := fiber.New(configurations.NewFiberConfig())
	app.Use(recover.New())

	fmt.Println("Loading CORS to allow from", os.Getenv("CORS"))
	fmt.Println("Allowing CORS to from", "*")
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CORS"),
		AllowHeaders: "*",
		AllowMethods: "GET, POST, PUT, OPTIONS",
	}))

	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Singapore",
	}))

	// Setup Routing
	meetingController.Route(app)

	port := os.Getenv("PORT")
	err := app.Listen("0.0.0.0:" + port)

	exceptions.PanicIfNeeded(err)
}
