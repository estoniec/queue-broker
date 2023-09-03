package app

import (
	"context"
	pb "github.com/estoniec/queue-broker/contracts/gen/go/queueBroker"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"queue-srvc/internal/config"
	v1 "queue-srvc/internal/controller/grpc/v1"
	"queue-srvc/internal/domain/queue/model"
	"queue-srvc/internal/domain/queue/service"
)

type App struct {
	config *config.Config
	server *v1.Server
}

func NewApp(cfg *config.Config) App {
	queue := model.NewQueue()
	queuePeoples := model.NewQueuePeoples()

	service := service.NewQueueService(queue, queuePeoples)
	server := v1.NewServer(service, pb.UnimplementedQueueServiceServer{})
	return App{
		config: cfg,
		server: server,
	}
}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return a.startGRPC()
	})

	return grp.Wait()
}

func (a *App) startGRPC() error {
	lis, err := net.Listen("tcp", a.config.Port)

	if err != nil {
		slog.Error("Failed to listing:", err)
	}

	slog.Info("Queue Svc on", a.config.Port)

	grpcServer := grpc.NewServer()

	pb.RegisterQueueServiceServer(grpcServer, a.server)

	return grpcServer.Serve(lis)
}
