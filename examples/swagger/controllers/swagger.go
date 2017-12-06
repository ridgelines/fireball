package controllers

import (
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/swagger/models"
	swagger "github.com/zpatrick/go-plugin-swagger"
)

type SwaggerController struct{}

func NewSwaggerController() *SwaggerController {
	return &SwaggerController{}
}

func (s *SwaggerController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/swagger.json",
			Handlers: fireball.Handlers{
				"GET": s.ServeSwaggerSpec,
			},
		},
	}

	return routes
}

func (s *SwaggerController) ServeSwaggerSpec(c *fireball.Context) (fireball.Response, error) {
	spec := swagger.Spec{
		SwaggerVersion: "2.0",
		Schemes:        []string{"http"},
		Info: &swagger.Info{
			Title:   "Swagger Example",
			Version: "0.0.1",
		},
		Definitions: map[string]swagger.Definition{
			"Movie": models.Movie{}.Definition(),
		},
		Tags: []swagger.Tag{
			{
				Name:        "Movies",
				Description: "Methods related to movies",
			},
		},
		Paths: map[string]swagger.Path{
			"/movies": map[string]swagger.Method{
				"get": {
					Summary: "List all Movies",
					Tags:    []string{"Movies"},
					Responses: map[string]swagger.Response{
						"200": {
							Description: "An array of movies",
							Schema:      swagger.NewObjectSliceSchema("Movie"),
						},
					},
				},
				"post": {
					Summary: "Add a Movie",
					Tags:    []string{"Movies"},
					Parameters: []swagger.Parameter{
						swagger.NewBodyParam("Movie", "Movie to add", true),
					},
					Responses: map[string]swagger.Response{
						"200": {
							Description: "The added movie",
							Schema:      swagger.NewObjectSchema("Movie"),
						},
					},
				},
			},
			"/movies/{title}": map[string]swagger.Method{
				"get": {
					Summary: "Describe a Movie",
					Tags:    []string{"Movies"},
					Parameters: []swagger.Parameter{
						swagger.NewStringPathParam("title", "Title of the movie to describe", true),
					},
					Responses: map[string]swagger.Response{
						"200": {
							Description: "The desired movie",
							Schema:      swagger.NewObjectSchema("Movie"),
						},
					},
				},
				"delete": {
					Summary: "Delete a Movie",
					Tags:    []string{"Movies"},
					Parameters: []swagger.Parameter{
						swagger.NewStringPathParam("title", "Title of the movie to delete", true),
					},
					Responses: map[string]swagger.Response{
						"200": {
							Description: "Success",
						},
					},
				},
			},
		},
	}

	return fireball.NewJSONResponse(200, spec)
}
