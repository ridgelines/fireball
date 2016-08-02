package main

import (
	"fmt"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/api/handlers"
	"github.com/zpatrick/fireball/examples/api/stores"
	"github.com/zpatrick/go-sdata/container"
	"net/http"
)

func main() {
	file := container.NewFileContainer("movies.json", nil)
	movieStore := stores.NewMovieStore(file)
	movieHandler := handlers.NewMovieHandler(movieStore)

	app := fireball.NewApp()
	// todo: app.Before = HTTPAuth()

	app.Routes = append(app.Routes, movieHandler.Routes()...)

	fmt.Println("Running on port 8000")
	http.ListenAndServe(":8000", app)
}
