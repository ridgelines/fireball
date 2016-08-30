package fireball

import (
	"net/http"
	"strings"
)

type Router interface {
	Match(*http.Request) (*RouteMatch, error)
}

type RouteMatch struct {
	Route         *Route
	Handler       Handler
	PathVariables map[string]string
}

type BasicRouter struct {
	Routes []*Route
	cache  map[string]*RouteMatch
}

func NewBasicRouter(routes []*Route) *BasicRouter {
	return &BasicRouter{
		Routes: routes,
		cache:  map[string]*RouteMatch{},
	}
}

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

	pathVariables, ok := r.matchPathVariables(route, req.URL.Path)
	if !ok {
		return nil
	}

	routeMatch := &RouteMatch{
		Route:         route,
		Handler:       handler,
		PathVariables: pathVariables,
	}

	return routeMatch
}

func (r *BasicRouter) matchPathVariables(route *Route, url string) (map[string]string, bool) {
	if url != "/" {
		url = strings.TrimSuffix(url, "/")
	}

	routeSections := strings.Split(route.Path, "/")
	urlSections := strings.Split(url, "/")

	if len(routeSections) != len(urlSections) {
		return nil, false
	}

	variables := map[string]string{}
	for i, routeSection := range routeSections {
		urlSection := urlSections[i]

		if strings.HasPrefix(routeSection, "{") && strings.HasSuffix(routeSection, "}") {
			key := routeSection[1 : len(routeSection)-1]
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
