package router

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/handlers"
	"al-mosso-api/internal/handlers/middlwares"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func startRoutes(router *fiber.App) {
	v1 := router.Group("/api/v1")

	{
		v1.Post("/contacts", handlers.MakeContactHandler)

		v1.Get("/chefs", handlers.GetChefsHandler)
		v1.Post("/chefs", handlers.InsertChefHandler)

		v1.Post("/foods", handlers.InsertFoodHandler)
		v1.Get("/foods", handlers.ListFoodHandler)
		v1.Get("/foods/:id", handlers.GetFoodById)

		v1.Post("/appointments", handlers.MakeAppointmentHandler)
		v1.Get("/appointments/confirm/:hash", handlers.ConfirmAppointmentHandler)
		v1.Post("/appointments/auth", handlers.AuthenticateAppointment)

		auth := v1.Group("" /*middleware*/, middlwares.CheckAuth)
		{
			auth.Use(jwtware.New(jwtware.Config{SigningKey: jwtware.SigningKey{Key: []byte(config.GetKey())}}))

			auth.Get("/appointments/:pin", handlers.GetAppointmentHandler)
			auth.Post("/appointments/add_food", handlers.AddFoodToAppointmentHandler)

		}

		//v1.Get("/appointments/:hash")

	}

}
