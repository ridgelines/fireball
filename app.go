package fireball

import (
	"fmt"
	"net/http"
	"sync"
)

type App struct {
	Routes []*Route
	// todo: Before Handler
	// todo: After Handler
	Error    func(http.ResponseWriter, error)
	NotFound func(http.ResponseWriter, *http.Request)
	once     sync.Once
}

func NewApp() *App {
	return &App{
		Error:    HandleError,
		NotFound: http.NotFound,
	}
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var match *Match
	for _, route := range a.Routes {
		if m := route.Match(r); m != nil {
			match = m
			break
		}
	}

	if match == nil {
		a.NotFound(w, r)
		return
	}

	c := &Context{
		PathVariables: match.Variables,
	}

	output, err := match.Handler(c)
	if err != nil {
		a.Error(w, err)
		return
	}

	if ok := tryWriteHeader(w, output); !ok {
		w.WriteHeader(http.StatusOK)
	}

	if ok := tryWriteResponse(w, output); !ok {
		data := fmt.Sprintf("%v", output)
		w.Write([]byte(data))
	}
}

func HandleError(w http.ResponseWriter, err error) {
	if ok := tryWriteHeader(w, err); !ok {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if ok := tryWriteResponse(w, err); !ok {
		w.Write([]byte(err.Error()))
	}
}
