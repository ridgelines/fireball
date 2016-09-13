package controllers

import (
	"encoding/json"
	"errors"
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/api/models"
	"github.com/zpatrick/fireball/examples/api/stores"
	"github.com/zpatrick/go-sdata/container"
	"net/http/httptest"
	"testing"
)

func newMovieStore(t *testing.T, movies []*models.Movie) *stores.MovieStore {
	container := container.NewMemoryContainer()
	store := stores.NewMovieStore(container)

	if err := store.Init(); err != nil {
		t.Fatal(err)
	}

	for _, movie := range movies {
		if err := store.Insert(movie).Execute(); err != nil {
			t.Fatal(err)
		}
	}

	return store
}

func jsonCompare(t *testing.T, response fireball.Response, expected interface{}) {
	w := httptest.NewRecorder()
	response.Write(w, nil)

	bytes, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	if v, want := w.Body.String(), string(bytes); v != want {
		t.Errorf("\nExpected: %s \nReceived: %s", want, v)
	}
}

func TestListMovies(t *testing.T) {
	movies := []*models.Movie{
		{
			ID:    "someID_1",
			Title: "someTitle_1",
			Year:  1,
		},
		{
			ID:    "someID_2",
			Title: "someTitle_2",
			Year:  2,
		},
	}

	store := newMovieStore(t, movies)
	controller := NewMovieController(store)

	response, err := controller.ListMovies(nil)
	if err != nil {
		t.Fatal(err)
	}

	jsonCompare(t, response, movies)
}

func TestListMoviesError(t *testing.T) {
	testContainer := container.NewTestContainer(nil)
	store := stores.NewMovieStore(testContainer)
	controller := NewMovieController(store)

	testContainer.SelectAllFunc = func(container.Container, string) (map[string][]byte, error) {
		return nil, errors.New("")
	}

	if _, err := controller.ListMovies(nil); err == nil {
		t.Fatal("Error was nil")
	}
}
