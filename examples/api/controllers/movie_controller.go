package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/api/models"
	"github.com/zpatrick/fireball/examples/api/stores"
	"math/rand"
)

type MovieController struct {
	Store *stores.MovieStore
}

func NewMovieController(store *stores.MovieStore) *MovieController {
	return &MovieController{
		Store: store,
	}
}

func (m *MovieController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/movies",
			Handlers: map[string]fireball.Handler{
				"GET":  m.ListMovies,
				"POST": m.CreateMovie,
			},
		},
		{
			Path: "/movies/{id}",
			Handlers: map[string]fireball.Handler{
				"GET":    m.GetMovie,
				"DELETE": m.DeleteMovie,
			},
		},
	}

	return routes
}

func (m *MovieController) ListMovies(c *fireball.Context) (fireball.Response, error) {
	movies, err := m.Store.SelectAll().Execute()
	if err != nil {
		return nil, err
	}

	return fireball.NewJSONResponse(200, movies, Headers)
}

func (m *MovieController) CreateMovie(c *fireball.Context) (fireball.Response, error) {
	var movie *models.Movie
	if err := json.NewDecoder(c.Request.Body).Decode(&movie); err != nil {
		return fireball.NewJSONError(400, err, Headers)
	}

	movie.ID = randomID(5)
	if err := m.Store.Insert(movie).Execute(); err != nil {
		return nil, err
	}

	return fireball.NewJSONResponse(200, movie, Headers)
}

func (m *MovieController) GetMovie(c *fireball.Context) (fireball.Response, error) {
	id := c.PathVariables["id"]

	movieIDMatch := func(m *models.Movie) bool {
		return m.ID == id
	}

	movie, err := m.Store.SelectAll().Where(movieIDMatch).FirstOrNil().Execute()
	if err != nil {
		return nil, err
	}

	if movie == nil {
		err := fmt.Errorf("Movie with id '%s' does not exist", id)
		return fireball.NewJSONError(400, err, Headers)
	}

	return fireball.NewJSONResponse(200, movie, Headers)
}

func (m *MovieController) DeleteMovie(c *fireball.Context) (fireball.Response, error) {
	id := c.PathVariables["id"]

	existed, err := m.Store.Delete(id).Execute()
	if err != nil {
		return nil, err
	}

	if !existed {
		err := fmt.Errorf("Movie with id '%s' does not exist", id)
		return fireball.NewJSONError(400, err, Headers)
	}

	return fireball.NewJSONResponse(200, nil, Headers)
}

const runes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomID(length int) string {
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = runes[rand.Intn(len(runes))]
	}

	return string(bytes)

}
