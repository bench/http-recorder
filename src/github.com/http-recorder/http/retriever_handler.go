package http

import (
	"encoding/json"
	"fmt"
	"github.com/http-recorder/entities"
	"github.com/http-recorder/fifo"
	nethttp "net/http"
	"strconv"
	"time"
)

func RetrieverHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	fmt.Println("[HTTP-RETRIEVER] (!) new client connection from", r.RemoteAddr)

	awaitingTime, headerNotAvailableOrMalformed := strconv.Atoi(r.Header.Get("Request-Timeout"))
	if headerNotAvailableOrMalformed != nil {
		awaitingTime = 3
	}

	awaitingChan := make(chan *entities.HttpRequest)
	stopChan := make(chan bool)
	go func() {

		cond := r.URL.Query().Get("pathcontains")
		if "" != cond { // Search by path
			fmt.Println("[HTTP-RETRIEVER] client asks for a request whose path contains", cond)
			request, err := fifo.FindByPath(cond)
			for err != nil {
				select {
				case <-stopChan:
					return
				default:
					request, err = fifo.FindByPath(cond)
					time.Sleep(1 * time.Second)
				}
			}
			awaitingChan <- request
		} else { // No search
			fmt.Println("[HTTP-RETRIEVER] client asks for any type of request")
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
		fmt.Println("[HTTP-RETRIEVER] sorry timeout reached, query returned no result, goodbye")
		w.WriteHeader(nethttp.StatusNotFound)
		stopChan <- true
	case request := <-awaitingChan:
		fmt.Println("[HTTP-RETRIEVER] return following request to client", request)
		json.NewEncoder(w).Encode(request)
	}

}
