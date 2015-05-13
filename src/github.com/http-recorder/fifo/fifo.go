package fifo

import (
	"fmt"
	"github.com/http-recorder/entities"
)

const (
	RequestFifoSize = 150
)

var requestFifo *Cache

func Init() {
	if onEvicted != nil {
		requestFifo = newWithOnEvicted(RequestFifoSize, onEvicted)
	}
	requestFifo = new(RequestFifoSize)
}

func onEvicted(value interface{}) {
	fmt.Println("memory is full, delete the following http request : ", value.(entities.HttpRequest))
}

func PersistRequest(hr *entities.HttpRequest) error {
	if requestFifo == nil {
		return fmt.Errorf("request list not initialized")
	}
	requestFifo.add(hr)
	return nil
}

func GetOldest() (*entities.HttpRequest, error) {
	oldestEntity := requestFifo.removeOldest()
	if oldestEntity == nil {
		return &entities.HttpRequest{}, fmt.Errorf("Queue is empty, noting to return")
	}
	return oldestEntity.Value.(*entities.HttpRequest), nil
}

func FindByPath(pattern string) (*entities.HttpRequest, error) {
	return searchRequestInFifo(MatcherFunc(containsPathMatcher), pattern)
}
