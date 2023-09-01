package model

import (
	"sync"
	"time"
)

type Queue struct {
	Mu    sync.RWMutex
	Queue map[string]map[int]string
}

type QueuePeoples struct {
	Mu           sync.RWMutex
	QueuePeoples map[string]map[int]time.Time
}

func NewQueue() *Queue {
	return &Queue{
		Mu:    sync.RWMutex{},
		Queue: make(map[string]map[int]string),
	}
}

func NewQueuePeoples() *QueuePeoples {
	return &QueuePeoples{
		Mu:           sync.RWMutex{},
		QueuePeoples: make(map[string]map[int]time.Time),
	}
}

//TODO: сделать создание айтема в очередь (типо ДТО)
