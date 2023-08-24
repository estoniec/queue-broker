package main

import (
	"api-gateway/internal/config"
	queue "api-gateway/internal/queueBroker"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func main() {
	c := config.GetConfig()

	app := fiber.New()

	queue.RegisterRoutes(app, c)

	slog.Error(app.Listen(c.Port).Error())
}
