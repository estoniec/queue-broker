package queue

import (
	"api-gateway/internal/config"
	"api-gateway/internal/queueBroker/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type ServiceClient struct {
	Client pb.QueueServiceClient
}

func InitServiceClient(c *config.Config) pb.QueueServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.BrokerSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error("Could not connect:", err)
	}

	return pb.NewQueueServiceClient(cc)
}
