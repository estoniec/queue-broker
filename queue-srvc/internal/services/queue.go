package services

import (
	"context"
	"queue-srvc/internal/pb"
	"sync"
	"time"
)

type Queue struct {
	Queue        *SafeMapQueue
	QueuePeoples *SafeMapQueuePeoples
	pb.UnimplementedQueueServiceServer
}

type SafeMapQueue struct {
	Mu    sync.RWMutex
	Queue map[string]map[int]string
}

type SafeMapQueuePeoples struct {
	Mu           sync.RWMutex
	QueuePeoples map[string]map[int]time.Time
}

// Queue очередь, где будут храниться элементы

func (h *Queue) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	if req.Category == "" {
		return &pb.PutResponse{
			Status: 404,
			Error:  "choose category",
		}, nil
	}
	if req.Item == "" {
		return &pb.PutResponse{
			Status: 404,
			Error:  "choose item",
		}, nil
	}
	h.Queue.Mu.Lock()
	defer h.Queue.Mu.Unlock()
	if h.Queue.Queue[req.Category] == nil {
		h.Queue.Queue[req.Category] = make(map[int]string)
	}
	h.Queue.Queue[req.Category][len(h.Queue.Queue[req.Category])] = req.Item
	return &pb.PutResponse{
		Status: 200,
		Error:  "",
	}, nil
}

func (h *Queue) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	if req.Category == "" {
		return &pb.GetResponse{
			Status: 404,
			Item:   "",
			Error:  "choose category",
		}, nil
	}
	var ID int
	time2 := time.Now()
	v := req.Timeout
	h.QueuePeoples.Mu.Lock()
	if h.QueuePeoples.QueuePeoples[req.Category] == nil {
		h.QueuePeoples.QueuePeoples[req.Category] = make(map[int]time.Time)
	}
	ID = len(h.QueuePeoples.QueuePeoples[req.Category])
	h.QueuePeoples.QueuePeoples[req.Category][ID] = time2
	if h.Queue.Queue[req.Category] == nil {
		if v == 0 {
			delete(h.QueuePeoples.QueuePeoples[req.Category], ID)
			h.QueuePeoples.Mu.Unlock()
			return &pb.GetResponse{
				Status: 404,
				Item:   "",
				Error:  "queue has 0 items in this category",
			}, nil
		}
		time := time.After(time.Duration(v) * time.Second)
		c := make(chan string)
		go h.waitQueue(req.Category, ID, c)
		for {
			select {
			case response := <-c:
				h.QueuePeoples.Mu.Unlock()
				return &pb.GetResponse{
					Status: 200,
					Item:   response,
					Error:  "",
				}, nil
			case <-time:
				delete(h.QueuePeoples.QueuePeoples[req.Category], ID)
				h.QueuePeoples.Mu.Unlock()
				return &pb.GetResponse{
					Status: 404,
					Item:   "",
					Error:  "timeout",
				}, nil
			}
		}
	}
	h.Queue.Mu.Lock()
	value, ok := h.Queue.Queue[req.Category][ID]
	h.Queue.Mu.Unlock()
	if !ok {
		if v == 0 {
			delete(h.QueuePeoples.QueuePeoples[req.Category], ID)
			h.QueuePeoples.Mu.Unlock()
			return &pb.GetResponse{
				Status: 404,
				Item:   "",
				Error:  "queue has 0 items in this category",
			}, nil
		}
		time := time.After(time.Duration(v) * time.Second)
		c := make(chan string)
		go h.waitQueue(req.Category, ID, c)
		for {
			select {
			case response := <-c:
				h.QueuePeoples.Mu.Unlock()
				return &pb.GetResponse{
					Status: 200,
					Item:   response,
					Error:  "",
				}, nil
			case <-time:
				delete(h.QueuePeoples.QueuePeoples[req.Category], ID)
				h.QueuePeoples.Mu.Unlock()
				return &pb.GetResponse{
					Status: 404,
					Item:   "",
					Error:  "timeout",
				}, nil
			}
		}
	}
	h.QueuePeoples.Mu.Unlock()
	return &pb.GetResponse{
		Status: 200,
		Item:   value,
		Error:  "",
	}, nil
}

func (h *Queue) waitQueue(path string, ID int, c chan string) {
	for {
		h.Queue.Mu.RLock()
		_, ok := h.Queue.Queue[path][ID]
		if !ok {
			h.Queue.Mu.RUnlock()
			continue
		}
		h.Queue.Mu.RUnlock()
		break
	}
	h.Queue.Mu.Lock()
	c <- h.Queue.Queue[path][ID]
	h.Queue.Mu.Unlock()
	return
}
