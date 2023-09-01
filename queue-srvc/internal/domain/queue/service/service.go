package service

import (
	"context"
	"queue-srvc/internal/domain/queue/dto"
	"queue-srvc/internal/domain/queue/model"
	"time"
)

type QueueService struct {
	Queue        *model.Queue
	QueuePeoples *model.QueuePeoples
}

// Queue очередь, где будут храниться элементы

func (h *QueueService) Put(ctx context.Context, input dto.PutQueueInput) (dto.PutQueueOutput, error) {
	if input.Category == "" {
		return dto.NewPutQueueOutput(404, "choose category"), nil
	}
	if input.Item == "" {
		return dto.NewPutQueueOutput(404, "choose item"), nil
	}
	h.Queue.Mu.Lock()
	defer h.Queue.Mu.Unlock()
	if h.Queue.Queue[input.Category] == nil {
		h.Queue.Queue[input.Category] = make(map[int]string)
	}
	h.Queue.Queue[input.Category][len(h.Queue.Queue[input.Category])] = input.Item
	return dto.NewPutQueueOutput(200, ""), nil
}

func (h *QueueService) Get(ctx context.Context, input dto.GetQueueInput) (dto.GetQueueOutput, error) {
	if input.Category == "" {
		return dto.NewGetQueueOutput(404, "", "choose category"), nil
	}
	var ID int
	time2 := time.Now()
	v := input.Timeout
	h.QueuePeoples.Mu.Lock()
	if h.QueuePeoples.QueuePeoples[input.Category] == nil {
		h.QueuePeoples.QueuePeoples[input.Category] = make(map[int]time.Time)
	}
	ID = len(h.QueuePeoples.QueuePeoples[input.Category])
	h.QueuePeoples.QueuePeoples[input.Category][ID] = time2
	if h.Queue.Queue[input.Category] == nil {
		if v == 0 {
			delete(h.QueuePeoples.QueuePeoples[input.Category], ID)
			h.QueuePeoples.Mu.Unlock()
			return dto.NewGetQueueOutput(404, "", "queue has 0 items in this category"), nil
		}
		time := time.After(time.Duration(v) * time.Second)
		c := make(chan string)
		go h.waitQueue(input.Category, ID, c)
		for {
			select {
			case response := <-c:
				h.QueuePeoples.Mu.Unlock()
				return dto.NewGetQueueOutput(200, response, ""), nil
			case <-time:
				delete(h.QueuePeoples.QueuePeoples[input.Category], ID)
				h.QueuePeoples.Mu.Unlock()
				return dto.NewGetQueueOutput(404, "", "timeout"), nil
			}
		}
	}
	h.Queue.Mu.Lock()
	value, ok := h.Queue.Queue[input.Category][ID]
	h.Queue.Mu.Unlock()
	if !ok {
		if v == 0 {
			delete(h.QueuePeoples.QueuePeoples[input.Category], ID)
			h.QueuePeoples.Mu.Unlock()
			return dto.NewGetQueueOutput(404, "", "queue has 0 items in this category"), nil
		}
		time := time.After(time.Duration(v) * time.Second)
		c := make(chan string)
		go h.waitQueue(input.Category, ID, c)
		for {
			select {
			case response := <-c:
				h.QueuePeoples.Mu.Unlock()
				return dto.NewGetQueueOutput(200, response, ""), nil
			case <-time:
				delete(h.QueuePeoples.QueuePeoples[input.Category], ID)
				h.QueuePeoples.Mu.Unlock()
				return dto.NewGetQueueOutput(404, "", "timeout"), nil
			}
		}
	}
	h.QueuePeoples.Mu.Unlock()
	return dto.NewGetQueueOutput(200, value, ""), nil
}

func (h *QueueService) waitQueue(path string, ID int, c chan string) {
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

func NewQueueService(queue *model.Queue, queuePeoples *model.QueuePeoples) *QueueService {
	return &QueueService{
		Queue:        queue,
		QueuePeoples: queuePeoples,
	}
}
