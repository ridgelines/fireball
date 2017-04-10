package main

import (
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/swagger/controllers"
	"log"
	"net/http"
)

const (
	SWAGGER_URL     = "/api/"
	SWAGGER_UI_PATH = "static/swagger-ui/dist"
)

func serveSwaggerUI(w http.ResponseWriter, r *http.Request) {
	dir := http.Dir(SWAGGER_UI_PATH)
	fileServer := http.FileServer(dir)
	http.StripPrefix(SWAGGER_URL, fileServer).ServeHTTP(w, r)
}

func main() {
	http.HandleFunc(SWAGGER_URL, serveSwaggerUI)

	routes := controllers.NewSwaggerController().Routes()
	routes = append(routes, controllers.NewMovieController().Routes()...)
	routes = fireball.Decorate(routes,
		fireball.LogDecorator())

	app := fireball.NewApp(routes)
	http.Handle("/", app)

	log.Println("Running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
