package router

import (
	"al-mosso-api/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func startRoutes(router *fiber.App) {
	v1 := router.Group("/api/v1")

	{

		v1.Get("/chefs", controller.GetChefsController)
		v1.Post("/chefs", controller.InsertChefController)

		v1.Post("/foods", controller.InsertFoodController)
		v1.Get("/foods", controller.ListFoodController)
	}

}
