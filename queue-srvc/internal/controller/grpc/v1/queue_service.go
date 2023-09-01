package v1

import (
	"context"
	pb "github.com/estoniec/queue-broker/contracts/gen/go/queueBroker"
	"log/slog"
)

// TODO: поработать над ДТО (по уркоам Артура)
func (s *Server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	input := NewPutQueueInput(req)

	res, err := s.service.Put(ctx, input)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewPutQueueOutput(res)

	return response, nil
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	input := NewGetQueueInput(req)

	res, err := s.service.Get(ctx, input)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewGetQueueOutput(res)

	return response, nil
}
