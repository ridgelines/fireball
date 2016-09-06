package main

import (
	"flag"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/rhyme/controllers"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	port := flag.String("p", "8000", "port to serve on")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
	go controllers.Init()

	rootController := controllers.NewRootController()
	app := fireball.NewApp(rootController.Routes())

	http.Handle("/", app)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Printf("Serving on port %s\n", *port)
	http.ListenAndServe(":"+*port, nil)
}
