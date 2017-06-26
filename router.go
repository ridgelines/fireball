package fireball

import (
	"net/http"
	"strings"
)

// Router is an interface that matches an *http.Request to a RouteMatch.
// If no matches are found, a nil RouteMatch should be returned.
type Router interface {
	Match(*http.Request) (*RouteMatch, error)
}

// RouterFunc is a function which implements the Router interface
type RouterFunc func(*http.Request) (*RouteMatch, error)

func (rf RouterFunc) Match(r *http.Request) (*RouteMatch, error) {
	return rf(r)
}

// BasicRouter attempts to match requests based on its Routes.
// This router supports variables in the URL by using ":variable" notation in URL sections.
// For example, the following are all valid Paths:
//  "/home"
//  "/movies/:id"
//  "/users/:userID/purchases/:purchaseID"
// Matched Path Variables can be retrieved in Handlers by the Context:
//  func Handler(c *Context) (Response, error) {
//      id := c.PathVariables["id"]
//      ...
//  }
type BasicRouter struct {
	Routes []*Route
	cache  map[string]*RouteMatch
}

// NewBasicRouter returns a new BasicRouter with the specified Routes
func NewBasicRouter(routes []*Route) *BasicRouter {
	return &BasicRouter{
		Routes: routes,
		cache:  map[string]*RouteMatch{},
	}
}

// Match attempts to match the *http.Request to a Route.
// Successful matches are cached for improved performance.
func (r *BasicRouter) Match(req *http.Request) (*RouteMatch, error) {
	if routeMatch, ok := r.cache[r.cacheKey(req)]; ok {
		return routeMatch, nil
	}

	for _, route := range r.Routes {
		if routeMatch := r.matchRoute(route, req); routeMatch != nil {
			r.cache[r.cacheKey(req)] = routeMatch
			return routeMatch, nil
		}
	}

	return nil, nil
}

func (r *BasicRouter) matchRoute(route *Route, req *http.Request) *RouteMatch {
	handler := route.Handlers[req.Method]
	if handler == nil {
		return nil
	}

	pathVariables, ok := r.matchPathVariables(route, req.URL.RawPath)
	if !ok {
		return nil
	}

	routeMatch := &RouteMatch{
		Handler:       handler,
		PathVariables: pathVariables,
	}

	return routeMatch
}

func (r *BasicRouter) matchPathVariables(route *Route, url string) (map[string]string, bool) {
	if url != "/" {
		url = strings.TrimSuffix(url, "/")
	}

	if route.Path != "/" {
		route.Path = strings.TrimSuffix(route.Path, "/")
	}

	routeSections := strings.Split(route.Path, "/")
	urlSections := strings.Split(url, "/")

	if len(routeSections) != len(urlSections) {
		return nil, false
	}

	variables := map[string]string{}
	for i, routeSection := range routeSections {
		urlSection := urlSections[i]

		if strings.HasPrefix(routeSection, ":") {
			key := routeSection[1:]
			variables[key] = urlSection
		} else if routeSection != urlSection {
			return nil, false
		}
	}

	return variables, true
}

func (r *BasicRouter) cacheKey(req *http.Request) string {
	return req.Method + req.URL.String()
}
