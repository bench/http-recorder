package http

import (
	"encoding/json"
	"fmt"
	"github.com/http-recorder/entities"
	"github.com/http-recorder/fifo"
	nethttp "net/http"
)

func RetrieverHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	fmt.Println("new client connects from", r.RemoteAddr)

	var request *entities.HttpRequest
	var err error

	// Handle only search on path
	if "" != r.URL.Query().Get("pathContains") {
		fmt.Println("client wants a request whose path contains", r.URL.Query().Get("pathContains"))
		request, err = fifo.FindByPath(r.URL.Query().Get("pathContains"))
	} else {
		fmt.Println("client wants any request")
		request, err = fifo.GetOldest()
	}

	if err != nil {
		fmt.Println("query returned no result, goodbye")
		w.WriteHeader(nethttp.StatusNotFound)
	} else {
		fmt.Println("'have a request for you", request)
		json.NewEncoder(w).Encode(request)
	}

}
