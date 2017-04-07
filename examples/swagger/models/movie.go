package models

import (
	"github.com/zpatrick/go-plugin-swagger"
)

type Movie struct {
	Title      string
	Year       int
	InTheaters bool
	Cast       []string
	Director   Contact
	Crew       []Contact
}

func (Movie) Definition() swagger.Definition {
	return swagger.Definition{
		Type: "object",
		Properties: map[string]swagger.Property{
			"title":       swagger.NewStringProperty(),
			"year":        swagger.NewIntProperty(),
			"in_theaters": swagger.NewBoolProperty(),
			"cast":        swagger.NewStringSliceProperty(),
			"directory":   swagger.NewObjectProperty("Contact"),
			"crew":        swagger.NewObjectSliceProperty("Contact"),
		},
	}
}
