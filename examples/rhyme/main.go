package main

import (
	"flag"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/rhyme/handlers"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	port := flag.String("p", "8000", "port to serve on")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
	go handlers.Init()

	rootHandler := handlers.NewRootHandler()
	app := fireball.NewApp()
	app.Routes = append(app.Routes, rootHandler.Routes()...)
	http.Handle("/", app)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Printf("Serving on port %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
