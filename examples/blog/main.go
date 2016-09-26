package main

import (
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/blog/controllers"
	"log"
	"net/http"
)

func main() {
	controller := controllers.NewRootController()
	routes := fireball.Decorate(controller.Routes(), fireball.BasicAuthDecorator("user", "pass"))
	app := fireball.NewApp(routes)
	app.Before = func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.String())
	}

	http.Handle("/", app)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe(":80", nil)
}
