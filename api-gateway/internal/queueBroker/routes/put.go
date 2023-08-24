package routes

import (
	"api-gateway/internal/queueBroker/pb"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Put(ctx *fiber.Ctx, c pb.QueueServiceClient) error {
	category := ctx.Path()
	v := ctx.Query("v", "")
	if v == "" {
		ctx.Status(404)
		ctx.Send(nil)
		return fmt.Errorf("can't add new value to queue without value")
	}

	_, err := c.Put(context.Background(), &pb.PutRequest{
		Category: category,
		Item:     v,
	})

	if err != nil {
		ctx.Status(404)
		ctx.Send(nil)
		return err
	}

	ctx.Status(200)
	ctx.Send(nil)
	return nil
}
