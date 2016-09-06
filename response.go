package fireball

import (
	"net/http"
)

type Response interface {
	Write(http.ResponseWriter, *http.Request)
}

type HTTPResponse struct {
	Status  int
	Body    []byte
	Headers map[string]string
}

func NewResponse(status int, body []byte, headers map[string]string) *HTTPResponse {
	return &HTTPResponse{
		Status:  status,
		Body:    body,
		Headers: headers,
	}
}

func (h *HTTPResponse) Write(w http.ResponseWriter, r *http.Request) {
	for key, val := range h.Headers {
		w.Header().Set(key, val)
	}

	w.WriteHeader(h.Status)
	w.Write(h.Body)
}
