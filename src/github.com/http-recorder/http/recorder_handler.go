package http

import (
	"encoding/json"
	"fmt"
	"github.com/http-recorder/entities"
	"github.com/http-recorder/fifo"
	nethttp "net/http"
)

func RecorderHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	fmt.Println("[HTTP-RECORDER] (!) New request received from", r.RemoteAddr)
	hr, err := entities.BuildHttpRequest(r)
	if err != nil {
		fmt.Println("[HTTP-RECORDER] Unable to process incoming request due to ", err)
		w.WriteHeader(nethttp.StatusBadRequest)
		return
	}
	fifo.PersistRequest(hr)
	bytes, _ := json.Marshal(hr)
	fmt.Println("[HTTP-RECORDER] Following stored with success :", string(bytes))
}
