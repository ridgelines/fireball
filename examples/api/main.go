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

var sessionStore = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	file := container.NewStringFileContainer("movies.json", nil)
	movieStore := stores.NewMovieStore(file)

	indexController := controllers.NewIndexController(sessionStore)
	movieController := controllers.NewMovieController(movieStore)

	routes := fireball.Decorate(
		append(movieController.Routes(), indexController.Routes()...),
		fireball.BasicAuthDecorator("user", "pass"),
		// todo: remove session decorator to different example (blog, put in username topright corner)
		fireball.SessionDecorator(sessionStore, time.Minute*1))

	app := fireball.NewApp(routes)
	app.ErrorHandler = controllers.JSONErrorHandler

	// Note that http://www.gorillatoolkit.org/pkg/sessions requires the use of context.ClearHandler:
	//  app := fireball.NewApp(routes)
	//  http.ListenAndServe(":8000", context.ClearHandler(app))

	fmt.Println("Running on port 8000")
	http.ListenAndServe(":8000", app)
}
