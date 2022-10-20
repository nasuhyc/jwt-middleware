package main

import (
	routes "jwt-middleware/backend/Routes"
	"jwt-middleware/backend/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.Connect()
	app := fiber.New()
	app.Use(logger.New())
	routes.Setup(app)
	app.Listen(":3001")
}
