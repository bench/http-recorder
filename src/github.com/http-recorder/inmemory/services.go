package inmemory

import (
	"fmt"
	"github.com/http-recorder/entities"
	"strings"
)

const (
	RequestFifoSize = 150
)

var requestFifo *Cache

func Init(onEvicted func(value interface{})) {
	if onEvicted != nil {
		requestFifo = newWithOnEvicted(RequestFifoSize, onEvicted)
	}
	requestFifo = new(RequestFifoSize)
}

func PersistRequest(r entities.HttpRequest) error {
	if requestFifo == nil {
		return fmt.Errorf("request list not initialized")
	}
	requestFifo.add(r)
	return nil
}

/* SEARCH FUNCTIONS */

func searchRequestPathContains(pattern string) (*entities.HttpRequest, error) {
	elements := requestFifo.getElements()
	for _, element := range elements {
		if strings.Contains(element.Value.(entities.HttpRequest).Path, pattern) {
			requestFifo.removeElement(element)
			hr := element.Value.(entities.HttpRequest)
			return &hr, nil
		}
	}
	return &entities.HttpRequest{}, fmt.Errorf("No request matching")
}
