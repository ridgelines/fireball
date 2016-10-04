package main

import (
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/api/controllers"
	"github.com/zpatrick/fireball/examples/api/stores"
	"github.com/zpatrick/go-sdata/container"
	"log"
	"net/http"
)

func main() {
	file := container.NewStringFileContainer("movies.json", nil)
	movieStore := stores.NewMovieStore(file)

	movieController := controllers.NewMovieController(movieStore)
	routes := fireball.Decorate(
		movieController.Routes(),
		fireball.BasicAuthDecorator("user", "pass"),
	)

	app := fireball.NewApp(routes)
	app.ErrorHandler = controllers.JSONErrorHandler

	log.Println("Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", app))
}
