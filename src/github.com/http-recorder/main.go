package main

import (
	"encoding/json"
	"fmt"
	"github.com/http-recorder/entities"
	"github.com/http-recorder/inmemory"
	"net/http"
)

const (
	RequestsBufferSize = 100
	RecorderServerPort = ":12345"
	ClientsServerPort  = ":23456"
)

func main() {
	fmt.Println("Starting http recorder...")
	inmemory.Init()
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
	inmemory.PersistRequest(hr)
	bytes, _ := json.Marshal(hr)
	fmt.Println("Request stored :", string(bytes))
}

func RequestProviderListener(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New client connects from", r.RemoteAddr)
	incomingRequests.Peek()
	json.NewEncoder(w).Encode(<-incomingRequests)
	w.WriteHeader(http.StatusOK)
}

func onEvictedRequest(value interface{}) {
	// FILL IT WITH YOUR NEED
	fmt.Println("Memory is full, remove oldest stored http request : ", value.(entities.HttpRequest))
}
