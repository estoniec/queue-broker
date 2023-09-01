package routes

import (
	"context"
	"fmt"
	pb "github.com/estoniec/queue-broker/contracts/gen/go/queueBroker"
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

	res, err := c.Put(context.Background(), &pb.PutRequest{
		Category: category,
		Item:     v,
	})

	if err != nil {
		fmt.Println(res)
		ctx.Status(int(res.Status))
		ctx.Send(nil)
		return err
	}

	if res.Error != "" {
		ctx.Status(int(res.Status))
		ctx.Send(nil)
		return fmt.Errorf(res.Error)
	}

	ctx.Status(int(res.Status))
	ctx.Send(nil)
	return nil
}
