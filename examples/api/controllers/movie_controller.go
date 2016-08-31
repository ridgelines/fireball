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

func (h *MovieController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		&fireball.Route{
			Path: "/movies",
			Handlers: map[string]fireball.Handler{
				"GET":  h.ListMovies,
				"POST": addAuth(h.CreateMovie),
			},
		},
		&fireball.Route{
			Path: "/movies/{id}",
			Handlers: map[string]fireball.Handler{
				"GET":    addAuth(h.GetMovie),
				"DELETE": addAuth(h.DeleteMovie),
			},
		},
	}

	return routes
}

func (h *MovieController) ListMovies(c *fireball.Context) (interface{}, error) {
	movies, err := h.Store.SelectAll().Execute()
	if err != nil {
		return nil, fireball.NewJSONError(500, err, nil)
	}

	return fireball.NewJSONResponse(200, movies, nil)
}

func (h *MovieController) CreateMovie(c *fireball.Context) (interface{}, error) {
	var movie models.Movie
	if err := json.NewDecoder(c.Request().Body).Decode(&movie); err != nil {
		return nil, fireball.NewJSONError(400, err, nil)
	}

	movie.ID = randomID(5)
	if err := h.Store.Insert(&movie).Execute(); err != nil {
		return nil, fireball.NewJSONError(500, err, nil)
	}

	return fireball.NewJSONResponse(200, movie, nil)
}

func (h *MovieController) GetMovie(c *fireball.Context) (interface{}, error) {
	id := c.PathVar("id")

	movieIDMatch := func(m *models.Movie) bool {
		return m.ID == id
	}

	movie, err := h.Store.SelectAll().Where(movieIDMatch).FirstOrNil().Execute()
	if err != nil {
		return nil, fireball.NewJSONError(500, err, nil)
	}

	if movie == nil {
		err := fmt.Errorf("Movie with id '%s' does not exist", id)
		return nil, fireball.NewJSONError(400, err, nil)
	}

	return fireball.NewJSONResponse(200, movie, nil)
}

func (h *MovieController) DeleteMovie(c *fireball.Context) (interface{}, error) {
	id := c.PathVar("id")

	existed, err := h.Store.Delete(id).Execute()
	if err != nil {
		return nil, fireball.NewJSONError(500, err, nil)
	}

	if !existed {
		err := fmt.Errorf("Movie with id '%s' does not exist", id)
		return nil, fireball.NewJSONError(400, err, nil)
	}

	return fireball.NewJSONResponse(200, nil, nil)
}

const runes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomID(length int) string {
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = runes[rand.Intn(len(runes))]
	}

	return string(bytes)

}
