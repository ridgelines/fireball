package fireball

import (
	"net/http"
	"strings"
)

type Router interface {
	Match(*Route, *http.Request) (*RouteMatch, error)
}

type RouteMatch struct {
	Route         *Route
	Handler       Handler
	PathVariables map[string]string
}

type BasicRouter struct {
	cache map[string]*RouteMatch
}

func NewBasicRouter() *BasicRouter {
	return &BasicRouter{
		cache: map[string]*RouteMatch{},
	}
}

func (r *BasicRouter) cacheKey(req *http.Request) string {
	return req.Method + req.URL.String()
}

func (r *BasicRouter) Match(route *Route, req *http.Request) (*RouteMatch, error) {
	if routeMatch, ok := r.cache[r.cacheKey(req)]; ok {
		return routeMatch, nil
	}

	handler := route.Handlers[req.Method]
	if handler == nil {
		return nil, nil
	}

	pathVariables, ok := r.matchPathVariables(route, req.URL.Path)
	if !ok {
		return nil, nil
	}

	routeMatch := &RouteMatch{
		Route:         route,
		Handler:       handler,
		PathVariables: pathVariables,
	}

	r.cache[r.cacheKey(req)] = routeMatch
	return routeMatch, nil
}

func (r *BasicRouter) matchPathVariables(route *Route, url string) (map[string]string, bool) {
	if url != "/" {
		url = strings.TrimSuffix(url, "/")
	}

	pathSections := strings.Split(route.Path, "/")
	urlSections := strings.Split(url, "/")

	if len(pathSections) != len(urlSections) {
		return nil, false
	}

	variables := map[string]string{}
	for i, pathSection := range pathSections {
		urlSection := urlSections[i]

		if strings.HasPrefix(pathSection, "{") && strings.HasSuffix(pathSection, "}") {
			key := pathSection[1 : len(pathSection)-1]
			variables[key] = urlSection
		} else if pathSection != urlSection {
			return nil, false
		}
	}

	return variables, true
}
