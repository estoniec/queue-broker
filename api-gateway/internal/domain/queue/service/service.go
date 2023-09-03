package service

import (
	"api-gateway/internal/domain/queue/dto"
	"context"
	pb "github.com/estoniec/queue-broker/contracts/gen/go/queueBroker"
)

type Service struct {
	client pb.QueueServiceClient
}

func NewService(client pb.QueueServiceClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) Get(context context.Context, input dto.GetQueueInput) (dto.GetQueueOutput, error) {
	res, err := s.client.Get(context, &pb.GetRequest{
		Category: input.Category,
		Timeout:  input.Timeout,
	})
	if err != nil {
		response := dto.NewGetQueueOutput(404, "", err.Error())
		return response, err
	}
	response := dto.NewGetQueueOutput(res.Status, res.Item, res.Error)
	return response, nil
}

func (s *Service) Put(context context.Context, input dto.PutQueueInput) (dto.PutQueueOutput, error) {
	res, err := s.client.Put(context, &pb.PutRequest{
		Category: input.Category,
		Item:     input.Item,
	})
	if err != nil {
		response := dto.NewPutQueueOutput(404, err.Error())
		return response, err
	}
	response := dto.NewPutQueueOutput(res.Status, res.Error)
	return response, nil
}
