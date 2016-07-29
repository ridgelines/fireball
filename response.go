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

type HTTPResponse struct {
	Data       []byte
	StatusCode int
}

func (r *HTTPResponse) Status() int {
	return r.StatusCode
}

func (r *HTTPResponse) Response() []byte {
	return r.Data
}

type HTMLResponse struct {
	*HTTPResponse
}

type JSONResponse struct {
	*HTTPResponse
}

func tryWriteHeader(w http.ResponseWriter, obj interface{}) bool {
	if obj, ok := obj.(Status); ok {
		w.WriteHeader(obj.Status())
		return true
	}

	return false
}

func tryWriteResponse(w http.ResponseWriter, obj interface{}) bool {
	if obj, ok := obj.(Response); ok {
		w.Write(obj.Response())
		return true
	}

	return false
}
