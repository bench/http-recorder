package fifo

import (
	"fmt"
	"github.com/http-recorder/entities"
	"github.com/http-recorder/log"
)

const (
	RequestFifoSize = 150
)

var requestFifo *Cache

func Init() {
	requestFifo = newWithOnEvicted(RequestFifoSize, onEvicted)
}

func onEvicted(value interface{}) {
	log.RecorderInfo("memory is full, delete the following http request : ", value.(*entities.HttpRequest))
}

func PersistRequest(hr *entities.HttpRequest) error {
	if requestFifo == nil {
		return fmt.Errorf("[HTTP-RECORDER] request list not initialized")
	}
	requestFifo.add(hr)
	return nil
}

func GetOldest() (*entities.HttpRequest, error) {
	oldestEntity := requestFifo.removeOldest()
	if oldestEntity == nil {
		return &entities.HttpRequest{}, fmt.Errorf("[HTTP-RETRIEVER] Queue is empty, nothing to return")
	}
	return oldestEntity.Value.(*entities.HttpRequest), nil
}
