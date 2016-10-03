package main

import (
	"fmt"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/api/controllers"
	"github.com/zpatrick/fireball/examples/api/stores"
	"github.com/zpatrick/go-sdata/container"
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

	fmt.Println("Running on port 8000")
	http.ListenAndServe(":8000", app)
}
