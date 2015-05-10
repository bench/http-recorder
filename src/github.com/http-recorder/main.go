package main

import (
	"encoding/json"
	"fmt"
	"github.com/http-recorder/entities"
	"github.com/http-recorder/lru"
	"net/http"
)

const (
	RequestsBufferSize = 100
	RecorderServerPort = ":12345"
	ClientsServerPort  = ":23456"
)

var requestQueuesByPath *lru.Cache

func main() {
	incomingRequests, _ = lru.New(128)
	fmt.Println("Starting http recorder...")
	go http.ListenAndServe(RecorderServerPort, http.HandlerFunc(RecorderListener))
	http.ListenAndServe(ClientsServerPort, http.HandlerFunc(RequestProviderListener))

}

func RecorderListener(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received new request to record from", r.RemoteAddr)
	hr, err := entities.BuildHttpRequest(r)
	if err != nil {
		fmt.Println("Unable to process incoming request due to ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bytes, _ := json.Marshal(hr)
	fmt.Println("Request stored :", string(bytes))
	incomingRequests.Add(hr)
}

func RequestProviderListener(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New client connects from", r.RemoteAddr)
	incomingRequests.Peek()
	json.NewEncoder(w).Encode(<-incomingRequests)
	w.WriteHeader(http.StatusOK)
}
