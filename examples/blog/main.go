package main

import (
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/blog/controllers"
	"log"
	"net/http"
)

func main() {
	controller := controllers.NewRootController()
	routes := fireball.Decorate(
		controller.Routes(),
		fireball.BasicAuthDecorator("user", "pass"),
		fireball.LogDecorator())

	app := fireball.NewApp(routes)
	http.Handle("/", app)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(":80", nil))
}
