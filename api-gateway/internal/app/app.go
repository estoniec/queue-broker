package app

import (
	"api-gateway/internal/config"
	v1 "api-gateway/internal/controller/http/v1"
	service2 "api-gateway/internal/domain/queue/service"
	"context"
	pb "github.com/estoniec/queue-broker/contracts/gen/go/queueBroker"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type App struct {
	app    *fiber.App
	config *config.Config
	client pb.QueueServiceClient
}

func NewApp(c *config.Config) *App {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.BrokerSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error("Could not connect:", err)
	}

	client := pb.NewQueueServiceClient(cc)

	app := fiber.New()

	service := service2.NewService(client)
	handler := v1.NewHandler(service, app, c)
	handler.Register()

	return &App{
		app:    app,
		config: c,
		client: pb.NewQueueServiceClient(cc),
	}
}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return a.startHTTP()
	})

	return grp.Wait()
}

func (a *App) startHTTP() error {
	return a.app.Listen(a.config.Port)
}
