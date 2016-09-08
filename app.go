package fireball

import (
	"net/http"
)

// App is the main structure of fireball applications.
// It can be invoked as an http.Handler
type App struct {
	// The After function is called after each request has completed
	After func(http.ResponseWriter, *http.Request)
	// The Before function is called before each request is routed
	Before func(http.ResponseWriter, *http.Request)
	// The ErrorHandler is called whenever a Handler returns a non-nil error
	ErrorHandler func(http.ResponseWriter, *http.Request, error)
	// The NotFoundHandler is called whenever the Router returns a nil RouteMatch
	NotFoundHandler func(http.ResponseWriter, *http.Request)
	// The template parser is passed into the Context
	Parser TemplateParser
	// The router is used to match a request to a Handler whenever a request is made
	Router Router
}

// NewApp returns a new App object with all of the default fields
func NewApp(routes []*Route) *App {
	return &App{
		After:           func(http.ResponseWriter, *http.Request) {},
		Before:          func(http.ResponseWriter, *http.Request) {},
		ErrorHandler:    DefaultErrorHandler,
		NotFoundHandler: http.NotFound,
		Parser:          NewGlobParser("views/", "*.html"),
		Router:          NewBasicRouter(routes),
	}
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Before(w, r)
	defer a.After(w, r)

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
		Meta:          map[string]interface{}{},
	}

	response, err := match.Handler(c)
	if err != nil {
		a.ErrorHandler(w, r, err)
		return
	}

	response.Write(w, r)
}
