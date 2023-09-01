package routes

import (
	"context"
	pb "github.com/estoniec/queue-broker/contracts/gen/go/queueBroker"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"strconv"
)

func Get(ctx *fiber.Ctx, c pb.QueueServiceClient) error {
	category := ctx.Path()
	timeout := ctx.Query("timeout", "")
	if timeout == "" {
		timeout = "0"
	}
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		ctx.Status(404)
		ctx.Send(nil)
		slog.Error(err.Error())
		return err
	}

	res, err := c.Get(context.Background(), &pb.GetRequest{
		Category: category,
		Timeout:  int64(timeoutInt),
	})

	if err != nil {
		ctx.Status(404)
		ctx.Send(nil)
		slog.Error(err.Error())
		return err
	}

	ctx.Status(200)
	ctx.Send([]byte(res.Item))
	return nil
}
