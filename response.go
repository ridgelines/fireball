package fireball

import (
	"log"
	"net/http"
)

// Response is an object that writes to an http.ResponseWriter
// A Response object implements the http.Handler interface
type Response interface {
	Write(http.ResponseWriter, *http.Request)
}

// ResponseFunc is a function which implements the Response interface
type ResponseFunc func(http.ResponseWriter, *http.Request)

func (rf ResponseFunc) Write(w http.ResponseWriter, r *http.Request) {
	rf(w, r)
}

// HTTPResponse objects write the specified status, headers, and body to
// a http.ResponseWriter
type HTTPResponse struct {
	Status  int
	Body    []byte
	Headers map[string]string
}

// NewResponse returns a new HTTPResponse with the specified status, body, and headers
func NewResponse(status int, body []byte, headers map[string]string) *HTTPResponse {
	return &HTTPResponse{
		Status:  status,
		Body:    body,
		Headers: headers,
	}
}

// Write will write the specified status, headers, and body to the http.ResponseWriter
func (h *HTTPResponse) Write(w http.ResponseWriter, r *http.Request) {
	for key, val := range h.Headers {
		w.Header().Set(key, val)
	}

	w.WriteHeader(h.Status)
	if _, err := w.Write(h.Body); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
