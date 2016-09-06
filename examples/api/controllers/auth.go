package controllers

import (
	"github.com/zpatrick/fireball"
)

func addAuth(handler fireball.Handler) fireball.Handler {
	return func(c *fireball.Context) (fireball.Response, error) {
		user, pass, ok := c.Request.BasicAuth()
		if ok && user == "user" && pass == "pass" {
			return handler(c)
		}

		headers := map[string]string{"WWW-Authenticate": "Basic realm=\"Restricted\""}
		response := fireball.NewResponse(401, []byte("401 Unauthorized\n"), headers)
		return response, nil
	}
}
