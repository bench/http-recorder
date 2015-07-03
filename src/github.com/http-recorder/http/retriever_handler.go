package http

import (
	"encoding/json"
	"github.com/http-recorder/entities"
	"github.com/http-recorder/fifo"
	"github.com/http-recorder/log"
	nethttp "net/http"
	"strconv"
	"time"
)

func RetrieverHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	log.RetrieverInfo("(!) new client connection from", r.RemoteAddr)

	awaitingTime, headerNotAvailableOrMalformed := strconv.Atoi(r.Header.Get("Request-Timeout"))
	if headerNotAvailableOrMalformed != nil {
		awaitingTime = 3
	}
	awaitingChan := make(chan *entities.HttpRequest)
	stopChan := make(chan bool)
	go func() {
		if len(r.URL.Query()) != 0 { // Search by path
			log.RetrieverInfo("client asks for a specific request", r.URL.Query())

			var request *entities.HttpRequest
			var err error

			for key, value := range r.URL.Query() {
				request, err = fifo.FindBy(key, value[0])
				break
			}
			// Handle long polling
			for err != nil {
				select {
				case <-stopChan:
					return
				default:
					for key, value := range r.URL.Query() {
						request, err = fifo.FindBy(key, value[0])
						break
					}
					time.Sleep(1 * time.Second)
				}
			}
			awaitingChan <- request
		} else { // No search
			log.RetrieverInfo("client asks for any type of request")
			request, err := fifo.GetOldest()
			for err != nil {
				select {
				case <-stopChan:
					return
				default:
					request, err = fifo.GetOldest()
					time.Sleep(1 * time.Second)
				}
			}
			awaitingChan <- request
		}
	}()

	select {
	case <-time.After(time.Second * time.Duration(awaitingTime)):
		log.RetrieverInfo("sorry timeout reached, query returned no result, goodbye")
		w.WriteHeader(nethttp.StatusNotFound)
		stopChan <- true
	case request := <-awaitingChan:
		log.RetrieverInfo("return following request to client", request)
		json.NewEncoder(w).Encode(request)
	}

}
