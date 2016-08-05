package handlers

import (
	"encoding/json"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/api/models"
	"github.com/zpatrick/fireball/examples/api/stores"
	"math/rand"
)

type MovieHandler struct {
	Store *stores.MovieStore
}

func NewMovieHandler(store *stores.MovieStore) *MovieHandler {
	return &MovieHandler{
		Store: store,
	}
}

func (h *MovieHandler) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		&fireball.Route{
			Path: "/movies",
			Get:  h.ListMovies,
			Post: h.CreateMovie,
		},
		&fireball.Route{
			Path:   "/movies/{id}",
			Get:    h.GetMovie,
			Delete: h.DeleteMovie,
		},
	}

	return routes
}

func (h *MovieHandler) ListMovies(c *fireball.Context) (interface{}, error) {
	movies, err := h.Store.SelectAll().Execute()
	if err != nil {
		return nil, c.JSONError(500, err)
	}

	return c.JSON(200, movies)
}

func (h *MovieHandler) CreateMovie(c *fireball.Context) (interface{}, error) {
	var movie models.Movie
	decoder := json.NewDecoder(c.Request().Body)

	if err := decoder.Decode(&movie); err != nil {
		return nil, c.JSONError(400, err)
	}

	movie.ID = randomID(10)

	if err := h.Store.Insert(&movie).Execute(); err != nil {
		return nil, c.JSONError(500, err)
	}

	return c.JSON(200, movie)
}

func (h *MovieHandler) GetMovie(c *fireball.Context) (interface{}, error) {
	id := c.PathVar("id")

	movieIDMatch := func(m *models.Movie) bool {
		return m.ID == id
	}

	movie, err := h.Store.SelectAll().Where(movieIDMatch).FirstOrNil().Execute()
	if err != nil {
		return nil, c.JSONError(500, err)
	}

	if movie == nil {
		return nil, c.JSONErrorf(400, "Movie with id '%s' does not exist", id)
	}

	return c.JSON(200, movie)
}

func (h *MovieHandler) DeleteMovie(c *fireball.Context) (interface{}, error) {
	id := c.PathVar("id")

	existed, err := h.Store.Delete(id).Execute()
	if err != nil {
		return nil, c.JSONError(500, err)
	}

	if !existed {
		return nil, c.JSONErrorf(400, "Movie with id '%s' does not exist", id)
	}

	return c.JSON(200, nil)
}

const runes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomID(length int) string {
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = runes[rand.Intn(len(runes))]
	}

	return string(bytes)

}
