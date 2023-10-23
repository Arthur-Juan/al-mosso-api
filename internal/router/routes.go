package router

import (
	"al-mosso-api/internal/handlers"
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

		v1.Post("/appointments", handlers.MakeAppointmentHandler)
		v1.Get("/appointments/confirm/:hash", handlers.ConfirmAppointmentHandler)
		//v1.Get("/appointments/:hash")

	}

}
