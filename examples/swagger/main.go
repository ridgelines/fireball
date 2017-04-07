package main

import (
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/swagger/controllers"
	"log"
	"net/http"
)

func main() {
	// serve swagger ui from /api
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	routes := controllers.NewSwaggerController().Routes()
	app := fireball.NewApp(routes)
	http.Handle("/", app)

	/*
	   fs := http.FileServer(http.Dir("static"))
	   http.Handle("/static/", http.StripPrefix("/static", fs))
	*/

	log.Println("Running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))

}
