package main

import (
	"github.com/zpatrick/fireball"
	"net/http"
)

func index(c *fireball.Context) (fireball.Response, error) {
	return fireball.NewResponse(200, []byte("Hello, World!"), nil), nil
}

func main() {
	indexRoute := &fireball.Route{
		Path: "/",
		Handlers: map[string]fireball.Handler{
			"GET": index,
		},
	}

	routes := []*fireball.Route{indexRoute}
	app := fireball.NewApp(routes)
	http.ListenAndServe(":8000", app)
}
