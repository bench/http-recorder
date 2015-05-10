package entities

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpRequest struct {
	Path    string              `json:"path"`
	Headers map[string][]string `json:"headers"`
	Method  string              `json:"method"`
	Body    string              `json:"body"`
}

func BuildHttpRequest(r *http.Request) (*HttpRequest, error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Http body unreadable")
	}
	return &HttpRequest{
		Path:    r.URL.Path,
		Headers: r.Header,
		Method:  r.Method,
		Body:    string(bytes),
	}, nil
}

func (hr *HttpRequest) String() string {
	return "PATH=" + hr.Path + "; METHOD=" + hr.Method + "; BODY=" + hr.Body
}
