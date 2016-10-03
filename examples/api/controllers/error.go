package controllers

import (
	"github.com/zpatrick/fireball"
	"log"
	"net/http"
)

func JSONErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	response, err := fireball.NewJSONError(500, err)
	if err != nil {
		log.Println(err)
		response := fireball.NewError(500, err, nil)
		response.Write(w, r)
		return
	}

	response.Write(w, r)
}
