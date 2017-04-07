package models

import (
	"github.com/zpatrick/go-plugin-swagger"
)

type Contact struct {
	Name  string
	Email string
}

func (Contact) Definition() swagger.Definition {
	return swagger.Definition{
		Type: "object",
		Properties: map[string]swagger.Property{
			"name":  swagger.NewStringProperty(),
			"email": swagger.NewStringProperty(),
		},
	}
}
