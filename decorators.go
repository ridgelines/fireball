package fireball

import (
	"log"
	"net/http"
)

// A Decorator wraps logic around a Handler
type Decorator func(Handler) Handler

// Decorate is a helper function that decorates each Handler in each Route with the given Decorators
func Decorate(routes []*Route, decorators ...Decorator) []*Route {
	decorated := make([]*Route, len(routes))

	for i, route := range routes {
		decorated[i] = &Route{
			Path:     route.Path,
			Handlers: map[string]Handler{},
		}

		for method, handler := range route.Handlers {
			for _, decorator := range decorators {
				handler = decorator(handler)
			}

			decorated[i].Handlers[method] = handler
		}
	}

	return decorated
}

// LogDecorator will print the method and url of each request
func LogDecorator() Decorator {
	return func(handler Handler) Handler {
		return func(c *Context) (Response, error) {
			log.Printf("%s %s \n", c.Request.Method, c.Request.URL.String())
			return handler(c)
		}
	}
}

// BasicAuthDecorator will add basic authentication using the specified username and password
func BasicAuthDecorator(username, password string) Decorator {
	return func(handler Handler) Handler {
		return func(c *Context) (Response, error) {
			user, pass, ok := c.Request.BasicAuth()
			if ok && user == username && pass == password {
				return handler(c)
			}

			headers := map[string]string{"WWW-Authenticate": "Basic realm=\"Restricted\""}
			response := NewResponse(401, []byte("401 Unauthorized\n"), headers)
			return response, nil
		}
	}
}

// HeaderResponseDecorator will add the specified headers to each response
func HeaderResponseDecorator(headers map[string]string) Decorator {
	return func(handler Handler) Handler {
		return func(c *Context) (Response, error) {
			response, err := handler(c)
			var wrappedResponse ResponseFunc = func(w http.ResponseWriter, r *http.Request) {
				for key, val := range headers {
					if v := w.Header().Get(key); v == "" {
						w.Header().Set(key, val)
					}
				}

				response.Write(w, r)
			}

			return wrappedResponse, err
		}
	}
}
