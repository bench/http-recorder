package http

import (
	"encoding/json"
	"github.com/http-recorder/entities"
	"github.com/http-recorder/fifo"
	"github.com/http-recorder/log"
	nethttp "net/http"
)

func RecorderHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	log.RecorderInfo("(!) new request received from", r.RemoteAddr)
	hr, err := entities.BuildHttpRequest(r)
	if err != nil {
		log.RecorderInfo("unable to process incoming request due to ", err)
		w.WriteHeader(nethttp.StatusBadRequest)
		return
	}
	fifo.PersistRequest(hr)
	bytes, _ := json.Marshal(hr)
	log.RecorderInfo("following stored with success :", string(bytes))
}
