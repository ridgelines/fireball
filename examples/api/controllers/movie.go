package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/api/models"
)

type MovieController struct {
	Movies map[string]models.Movie
}

func NewMovieController() *MovieController {
	return &MovieController{
		Movies: map[string]models.Movie{},
	}
}

func (m *MovieController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/movies",
			Handlers: fireball.Handlers{
				"GET":  m.ListMovies,
				"POST": m.AddMovie,
			},
		},
		{
			Path: "/movies/:title",
			Handlers: fireball.Handlers{
				"GET":    m.GetMovie,
				"DELETE": m.DeleteMovie,
			},
		},
	}

	return routes
}

func (m *MovieController) ListMovies(c *fireball.Context) (fireball.Response, error) {
	movies := []models.Movie{}
	for _, movie := range m.Movies {
		movies = append(movies, movie)
	}

	return fireball.NewJSONResponse(200, movies)
}

func (m *MovieController) AddMovie(c *fireball.Context) (fireball.Response, error) {
	var movie models.Movie
	if err := json.NewDecoder(c.Request.Body).Decode(&movie); err != nil {
		return nil, err
	}

	m.Movies[movie.Title] = movie
	return fireball.NewJSONResponse(200, movie)
}

func (m *MovieController) GetMovie(c *fireball.Context) (fireball.Response, error) {
	title := c.PathVariables["title"]
	movie, ok := m.Movies[title]
	if !ok {
		return nil, fmt.Errorf("Movie with title '%s' does not exist", title)
	}

	return fireball.NewJSONResponse(200, movie)
}

func (m *MovieController) DeleteMovie(c *fireball.Context) (fireball.Response, error) {
	title := c.PathVariables["title"]
	if _, ok := m.Movies[title]; !ok {
		return nil, fmt.Errorf("Movie with title '%s' does not exist", title)
	}

	delete(m.Movies, title)
	return fireball.NewJSONResponse(200, nil)
}
