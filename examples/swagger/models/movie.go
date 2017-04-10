package models

import (
	"github.com/zpatrick/go-plugin-swagger"
)

type Movie struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func (m Movie) Definition() swagger.Definition {
	return swagger.Definition{
		Type: "object",
		Properties: map[string]swagger.Property{
			"title": swagger.NewStringProperty(),
			"year":  swagger.NewIntProperty(),
		},
	}
}
