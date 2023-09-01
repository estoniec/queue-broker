package v1

import (
	pb "github.com/estoniec/queue-broker/contracts/gen/go/queueBroker"
	"queue-srvc/internal/domain/queue/service"
)

type Server struct {
	service *service.QueueService
	pb.UnimplementedQueueServiceServer
}

func NewServer(
	service *service.QueueService,
	srv pb.UnimplementedQueueServiceServer,
) *Server {
	return &Server{
		service:                         service,
		UnimplementedQueueServiceServer: srv,
	}
}
