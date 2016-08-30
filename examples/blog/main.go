package main

import (
	"fmt"
	"github.com/zpatrick/fireball"
	"net/http"
)

func main() {
	routes := []*fireball.Route{
		&fireball.Route{
			Path: "/",
			Handlers: map[string]fireball.Handler{
				"GET": func(c *fireball.Context) (interface{}, error) {
					return "Hello, World!", nil
				},
			},
		},
	}

	app := fireball.NewApp(routes)

	fmt.Println("Running on port 8000")
	http.ListenAndServe(":8000", app)
}
