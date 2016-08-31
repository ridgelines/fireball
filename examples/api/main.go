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

	app := fireball.NewApp(movieController.Routes())

	fmt.Println("Running on port 8000")
	http.ListenAndServe(":8000", app)
}
