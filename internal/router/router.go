package router

import "github.com/gofiber/fiber/v2"

func Initialize() {
	router := fiber.New()
	router.Static("/uploads", "./uploads")
	startRoutes(router)
	err := router.Listen(":8080")
	if err != nil {
		return
	}
}
