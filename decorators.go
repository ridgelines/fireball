package fireball

import (
	"github.com/gorilla/sessions"
	"log"
	"time"
)

type Decorator func(Handler) Handler

func Decorate(routes []*Route, decorators ...Decorator) []*Route {
	for _, decorator := range decorators {
		for _, route := range routes {
			for method, handler := range route.Handlers {
				route.Handlers[method] = decorator(handler)
			}
		}
	}

	return routes
}

func BasicAuthDecorator(username, password string) Decorator {
	return func(handler Handler) Handler {
		return func(c *Context) (interface{}, error) {
			user, pass, ok := c.Request.BasicAuth()
			if ok && user == username && pass == password {
				return handler(c)
			}

			headers := map[string]string{"WWW-Authenticate": "Basic realm=\"Restricted\""}
			response := NewHTTPResponse(401, []byte("401 Unauthorized\n"), headers)
			return response, nil
		}
	}
}

func LogDecorator() Decorator {
	return func(handler Handler) Handler {
		return func(c *Context) (interface{}, error) {
			log.Printf("%s %s\n", c.Request.Method, c.Request.URL.String())
			return handler(c)
		}
	}
}

// todo: from gorilla documentation http://www.gorillatoolkit.org/pkg/sessions
// need use context.ClearHandler: http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux)))
func SessionDecorator(store sessions.Store, expiration time.Duration) Decorator {
	return func(handler Handler) Handler {
		return func(c *Context) (interface{}, error) {
			session, err := store.Get(c.Request, "session")
			if err != nil {
				return nil, err
			}

			session.Options.MaxAge = int(expiration.Seconds())
			c.Meta["session"] = session
			defer session.Save(c.Request, c.Writer)
			return handler(c)
		}
	}
}
