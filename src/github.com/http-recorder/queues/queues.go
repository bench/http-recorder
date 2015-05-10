package queues

import (
	"github.com/http-recorder/entities"
)

var requestQueuesByPath *Cache = New(15)

func PersistRequest(path string, r entities.HttpRequest) {
	ensureQueueExists(path)
	requestQueuesByPath.Get(path)
}

func ensureQueueExists(path string) {
	if requestQueuesByPath.Contains(path) {
		return
	}
	requestQueuesByPath.Add(make(chan entities.HttpRequest, 30))
}
