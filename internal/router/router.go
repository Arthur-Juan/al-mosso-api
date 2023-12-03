package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func Initialize() {
	router := fiber.New()
	router.Static("/uploads", "./uploads")
	router.Static("/uploads/*", "./uploads/*")
	router.Static("/docs", "./docs")

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "*",
	}))

	router.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "http://localhost:8080/docs/swagger.json",
	}))
	startRoutes(router)
	err := router.Listen(":8080")
	if err != nil {
		return
	}
}
