package handlers

import (
	"github.com/zpatrick/fireball"
)

func addAuth(handler fireball.Handler) fireball.Handler {
	return func(c *fireball.Context) (interface{}, error) {
		user, pass, _ := c.Request().BasicAuth()

		if user == "user" && pass == "pass" {
			return handler(c)
		}

		headers := map[string]string{"WWW-Authenticate": "Basic realm=\"Restricted\""}
		response := fireball.NewHTTPResponse(401, []byte("401 Unauthorized\n"), headers)
		return response, nil
	}
}
