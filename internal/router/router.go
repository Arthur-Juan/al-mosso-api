package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Initialize() {
	router := fiber.New()
	router.Static("/uploads", "./uploads")
	router.Static("/uploads/*", "./uploads/*")

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "*",
	}))
	startRoutes(router)
	err := router.Listen(":8080")
	if err != nil {
		return
	}
}
