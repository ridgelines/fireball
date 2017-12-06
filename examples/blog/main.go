package main

import (
	"log"
	"net/http"

	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/blog/controllers"
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

	log.Println("Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
