package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/api/controllers"
	"github.com/zpatrick/fireball/examples/api/stores"
	"github.com/zpatrick/go-sdata/container"
	"net/http"
	"time"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	file := container.NewStringFileContainer("movies.json", nil)
	movieStore := stores.NewMovieStore(file)

	indexController := controllers.NewIndexController(store)
	movieController := controllers.NewMovieController(movieStore)

	routes := fireball.Decorate(
		append(movieController.Routes(), indexController.Routes()...),
		fireball.LogDecorator(),
		fireball.BasicAuthDecorator("user", "pass"),
		fireball.SessionDecorator(store, time.Minute*1))

	app := fireball.NewApp(routes)
	app.ErrorHandler = controllers.JSONErrorHandler

	fmt.Println("Running on port 8000")
	http.ListenAndServe(":8000", app)
}
