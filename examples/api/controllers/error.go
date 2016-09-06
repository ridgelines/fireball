package controllers

import (
	"github.com/zpatrick/fireball"
	"log"
	"net/http"
)

func JSONErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	response, err := fireball.NewJSONError(500, err, fireball.JSONHeaders)
	if err != nil {
		log.Println(err)
		response := fireball.NewError(500, err, fireball.JSONHeaders)
		response.Write(w, r)
		return
	}

	response.Write(w, r)
}
