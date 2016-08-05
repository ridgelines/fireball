package fireball

import (
	"net/http"
)

type Status interface {
	Status() int
}

type Response interface {
	Response() []byte
}

type Header interface {
	Header() map[string]string
}

type HTTPResponse struct {
	Data       []byte
	StatusCode int
	Headers    map[string]string
}

func (r *HTTPResponse) Status() int {
	return r.StatusCode
}

func (r *HTTPResponse) Response() []byte {
	return r.Data
}

func (r *HTTPResponse) Header() map[string]string {
	return r.Headers
}

type HTMLResponse struct {
	*HTTPResponse
}

type JSONResponse struct {
	*HTTPResponse
}

func tryWriteHeader(w http.ResponseWriter, obj interface{}) bool {
	var didWrite bool

	if obj, ok := obj.(Header); ok {
		for key, val := range obj.Header() {
			w.Header().Set(key, val)
		}

		didWrite = true
	}

	if obj, ok := obj.(Status); ok {
		w.WriteHeader(obj.Status())
		didWrite = true
	}

	return didWrite
}

func tryWriteResponse(w http.ResponseWriter, obj interface{}) bool {
	if obj, ok := obj.(Response); ok {
		w.Write(obj.Response())
		return true
	}

	return false
}
