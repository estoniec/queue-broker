package v1

import (
	pb "github.com/estoniec/queue-broker/contracts/gen/go/queueBroker"
	"queue-srvc/internal/domain/queue/dto"
)

func NewPutQueueInput(req *pb.PutRequest) dto.PutQueueInput {
	return dto.NewPutQueueInput(req.Category, req.Item)
}

func NewPutQueueOutput(out dto.PutQueueOutput) *pb.PutResponse {
	return &pb.PutResponse{
		Status: out.Status,
		Error:  out.Error,
	}
}

func NewGetQueueInput(req *pb.GetRequest) dto.GetQueueInput {
	return dto.NewGetQueueInput(req.Timeout, req.Category)
}

func NewGetQueueOutput(out dto.GetQueueOutput) *pb.GetResponse {
	return &pb.GetResponse{
		Status: out.Status,
		Item:   out.Item,
		Error:  out.Error,
	}
}
