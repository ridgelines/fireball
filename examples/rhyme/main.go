package main

import (
	"fmt"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/rhyme/handlers"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	rootHandler := handlers.NewRootHandler()

	app := fireball.NewApp()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	//fs := http.FileServer(http.Dir("./static"))
	//http.Handle("/static", fs) //http.StripPrefix("/static", fs))

	app.Routes = append(app.Routes, rootHandler.Routes()...)

	fmt.Println("Running on port 8000")
	//http.ListenAndServe(":8000", app)
	http.Handle("/", app)
	http.ListenAndServe(":8000", nil)
}
