package fireball

import (
	"github.com/gorilla/sessions"
	"net/http"
	"time"
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
			decorated[i].Handlers[method] = handler

			for _, decorator := range decorators {
				decorated[i].Handlers[method] = decorator(decorated[i].Handlers[method])
			}
		}
	}

	return decorated
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

// SessionDecorator will manage a *gorilla.Session object.
// The session can be accessed by the "session" key in the Context.Meta field.
//
// Note that http://www.gorillatoolkit.org/pkg/sessions requires the use of context.ClearHandler:
//  app := fireball.NewApp(routes)
//  http.ListenAndServe(":8000", context.ClearHandler(app))
func SessionDecorator(store sessions.Store, expiration time.Duration) Decorator {
	return func(handler Handler) Handler {
		return func(c *Context) (Response, error) {
			session, err := store.Get(c.Request, "session")
			if err != nil {
				return nil, err
			}

			session.Options.MaxAge = int(expiration.Seconds())
			c.Meta["session"] = session

			response, err := handler(c)
			var wrappedResponse ResponseFunc = func(w http.ResponseWriter, r *http.Request) {
				session.Save(r, w)
				response.Write(w, r)
			}

			return wrappedResponse, err
		}
	}
}
