package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
	"queue-srvc/internal/config"
	"queue-srvc/internal/pb"
	"queue-srvc/internal/services"
	"sync"
	"time"
)

func main() {
	c := config.GetConfig()

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		slog.Error("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Queue{
		Queue: &services.SafeMapQueue{
			Mu:    sync.RWMutex{},
			Queue: make(map[string]map[int]string),
		},
		QueuePeoples: &services.SafeMapQueuePeoples{
			Mu:           sync.RWMutex{},
			QueuePeoples: make(map[string]map[int]time.Time),
		},
	}

	grpcServer := grpc.NewServer()

	pb.RegisterQueueServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
