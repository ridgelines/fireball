package fireball

import (
	"net/http"
	"sync"
)

type App struct {
	Parser          TemplateParser
	Router          Router
	ErrorHandler    func(http.ResponseWriter, *http.Request, error)
	NotFoundHandler func(http.ResponseWriter, *http.Request)
	once            sync.Once
}

func NewApp(routes []*Route) *App {
	parser := &GlobParser{
		Root: "views/",
		Glob: "*.html",
	}

	return &App{
		ErrorHandler:    DefaultErrorHandler,
		NotFoundHandler: http.NotFound,
		Router:          NewBasicRouter(routes),
		Parser:          parser,
	}
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	match, err := a.Router.Match(r)
	if err != nil {
		a.ErrorHandler(w, r, err)
		return
	}

	if match == nil {
		a.NotFoundHandler(w, r)
		return
	}

	c := &Context{
		PathVariables: match.PathVariables,
		Parser:        a.Parser,
		Request:       r,
		Writer:        w,
		Meta:          map[string]interface{}{},
	}

	response, err := match.Handler(c)
	if err != nil {
		a.ErrorHandler(w, r, err)
		return
	}

	response.Write(w, r)
}
