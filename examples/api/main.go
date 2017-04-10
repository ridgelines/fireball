package main

import (
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/api/controllers"
	"log"
	"net/http"
)

func main() {
	routes := controllers.NewMovieController().Routes()
	routes = fireball.Decorate(routes,
		fireball.BasicAuthDecorator("user", "pass"),
		fireball.LogDecorator())

	app := fireball.NewApp(routes)

	log.Println("Running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", app))
}
