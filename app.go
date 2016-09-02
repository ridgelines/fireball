package fireball

import (
	"fmt"
	"net/http"
	"sync"
)

type App struct {
	Parser   TemplateParser
	Router   Router
	Error    func(http.ResponseWriter, error)
	NotFound func(http.ResponseWriter, *http.Request)
	once     sync.Once
}

func NewApp(routes []*Route) *App {
	parser := &GlobParser{
		Root: "views/",
		Glob: "*.html",
	}

	return &App{
		Error:    HandleError,
		NotFound: http.NotFound,
		Router:   NewBasicRouter(routes),
		Parser:   parser,
	}
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	match, err := a.Router.Match(r)
	if err != nil {
		a.Error(w, err)
		return
	}

	if match == nil {
		a.NotFound(w, r)
		return
	}

	c := &Context{
		PathVariables: match.PathVariables,
		Parser:        a.Parser,
		Request:       r,
		Writer:        w,
		Meta:          map[string]interface{}{},
	}

	// todo: what if output is a type of fireball.Redirect?
	// would that work at all?
	// type Redirect func() *Route
	// Handler(c){ return c.Redirect(this.Index), nil }
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
