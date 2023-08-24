package queue

import (
	"api-gateway/internal/config"
	"api-gateway/internal/queueBroker/routes"
	"github.com/gofiber/fiber/v2"
)

var handler = func(c *fiber.Ctx) error { return nil }

func RegisterRoutes(r *fiber.App, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/")
	routes.Put("/:category", svc.Put)
	routes.Get("/:category", svc.Get)

	return svc
}

func (svc *ServiceClient) Put(ctx *fiber.Ctx) error {
	return routes.Put(ctx, svc.Client)
}

func (svc *ServiceClient) Get(ctx *fiber.Ctx) error {
	return routes.Get(ctx, svc.Client)
}
