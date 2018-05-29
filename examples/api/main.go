package main

import (
	"log"
	"net/http"

	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/api/controllers"
)

func main() {
	routes := controllers.NewMovieController().Routes()
	routes = fireball.Decorate(routes,
		fireball.BasicAuthDecorator("user", "pass"),
		fireball.LogDecorator())

	app := fireball.NewApp(routes)
	app.ErrorHandler = controllers.JSONErrorHandler

	log.Println("Running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", app))
}
