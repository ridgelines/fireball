package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/go-plugin-swagger"
)

type SwaggerController struct {
}

func NewSwaggerController() *SwaggerController {
	return &SwaggerController{}
}

func (s *SwaggerController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/api",
			Handlers: fireball.Handlers{
				"GET": s.swaggerUIRedirect,
			},
		},
		{
			Path: "/api/spec.json",
			Handlers: fireball.Handlers{
				"GET": s.getSwaggerSpec,
			},
		},
	}

	return routes
}

func (s *SwaggerController) swaggerUIRedirect(c *fireball.Context) (fireball.Response, error) {
	redirectPath := fmt.Sprintf("/static/swagger?url=%s/docs.json", c.Request.URL.String())
	return fireball.Redirect(301, redirectPath), nil
}

func (s *SwaggerController) getSwaggerSpec(c *fireball.Context) (fireball.Response, error) {
	spec := swagger.Spec{
		SwaggerVersion: "2.0",
		Schemes:        []string{"https"},
		Info: swagger.Info{
			Title:          "title",
			Version:        "1.0.0",
			TermsOfService: "tos",
			Contact: swagger.Contact{
				Name:  "First Last",
				Email: "first_last@email.com",
				URL:   "http://url.domain.com",
			},
			License: swagger.License{
				Name: "MIT",
				URL:  "http://url.domain.com",
			},
		},
		Tags: []swagger.Tag{
			{
				Name:         "movies",
				Description:  "api calls about movies",
				ExternalDocs: swagger.ExternalDocs{},
			},
		},
		Paths: map[string]swagger.Path{
			"/movies": map[string]swagger.Method{
				"post": {
					Summary:     "Create a Movie",
					Description: "SLorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
					Parameters: []swagger.Parameter{
						swagger.NewIntPathParam("count", "some description", true),
						swagger.NewStringPathParam("name", "some description", true),
						swagger.NewBodyParam(true, models.Movie{})
					},
					Responses: map[string]swagger.Response{
						"200": {
							Description: "Successful Operation",
						},
						"400": {
							Description: "Invalid Request",
						},
					},
				},
			},
		},
	}

	bytes, err := json.MarshalIndent(spec, "", "    ")
	if err != nil {
		return nil, err
	}

	return fireball.NewResponse(200, bytes, fireball.JSONHeaders), nil
}
