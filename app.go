package fireball

import (
	"fmt"
	"net/http"
	"sync"
)

type App struct {
	Parser TemplateParser
	Routes []*Route
	// todo: Before Handler
	// todo: After Handler
	Error    func(http.ResponseWriter, error)
	NotFound func(http.ResponseWriter, *http.Request)
	once     sync.Once
}

func NewApp() *App {
	parser := &GlobParser{
		Root: "views/",
		Glob: "*.html",
	}

	return &App{
		Error:    HandleError,
		NotFound: http.NotFound,
		Parser:   parser,
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
		Parser:        a.Parser,
		request:       r,
	}

	output, err := match.Handler(c)
	if err != nil {
		a.Error(w, err)
		return
	}

	if ok := TryWriteHeader(w, output); !ok {
		w.WriteHeader(http.StatusOK)
	}

	if ok := TryWriteResponse(w, output); !ok {
		data := fmt.Sprintf("%v", output)
		w.Write([]byte(data))
	}
}

func HandleError(w http.ResponseWriter, err error) {
	if ok := TryWriteHeader(w, err); !ok {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if ok := TryWriteResponse(w, err); !ok {
		w.Write([]byte(err.Error()))
	}
}

func TryWriteHeader(w http.ResponseWriter, obj interface{}) bool {
	var didWrite bool

	if obj, ok := obj.(Headers); ok {
		for key, val := range obj.Headers() {
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

func TryWriteResponse(w http.ResponseWriter, obj interface{}) bool {
	if obj, ok := obj.(Body); ok {
		w.Write(obj.Body())
		return true
	}

	return false
}
