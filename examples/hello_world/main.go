package main

import (
	"github.com/zpatrick/fireball"
	"log"
	"net/http"
)

func index(c *fireball.Context) (fireball.Response, error) {
	return fireball.NewResponse(200, []byte("Hello, World!"), nil), nil
}

func main() {
	indexRoute := &fireball.Route{
		Path: "/",
		Handlers: fireball.Handlers{
			"GET": index,
		},
	}

	routes := []*fireball.Route{indexRoute}
	app := fireball.NewApp(routes)

	log.Println("Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", app))
}
