package main

import (
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/blog/controllers"
	"net/http"
)

func main() {
	controller := controllers.NewIndexController()
	routes := fireball.Decorate(controller.Routes(), fireball.LogDecorator())
	app := fireball.NewApp(routes)

	http.Handle("/", app)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe(":8000", nil)
}
