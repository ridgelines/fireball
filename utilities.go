package fireball

import (
	"net/http"
)

// Redirect wraps http.Redirect in a ResponseFunc
func Redirect(status int, url string) Response {
	return ResponseFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, status)
	})
}

// EnableCORS decorates each route by adding CORS headers to each response
// An OPTIONS Handler is added to each route if one doesn't already exist
func EnableCORS(routes []*Route) []*Route {
	decorated := Decorate(routes, HeaderResponseDecorator(CORSHeaders))

	for _, route := range decorated {
		if _, exists := route.Handlers["OPTIONS"]; !exists {
			route.Handlers["OPTIONS"] = func(c *Context) (Response, error) {
				return NewResponse(200, nil, CORSHeaders), nil
			}
		}
	}

	return decorated
}
