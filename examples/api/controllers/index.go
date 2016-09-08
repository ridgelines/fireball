package controllers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/zpatrick/fireball"
	"net/http"
)

var Headers = fireball.JSONHeaders

type IndexController struct {
	Store sessions.Store
}

func NewIndexController(store sessions.Store) *IndexController {
	return &IndexController{
		Store: store,
	}
}

func (i *IndexController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/",
			Handlers: map[string]fireball.Handler{
				"GET": i.Redirect,
			},
		},
	}

	return routes
}

// example not using decorator
func (h *IndexController) IndexWithoutDecorator(c *fireball.Context) (fireball.Response, error) {
	session, err := h.Store.Get(c.Request, "session")
	if err != nil {
		return nil, err
	}

	count, ok := session.Values["count"].(int)
	if !ok {
		count = 0
	}

	session.Values["count"] = count + 1
	body := fmt.Sprintf("You have visited this page %d time(s)", session.Values["count"])
	response := fireball.NewResponse(200, []byte(body), nil)
	return saveSession(response, session), nil
}

func saveSession(response fireball.Response, session *sessions.Session) fireball.Response {
	var wrappedResponse fireball.ResponseFunc = func(w http.ResponseWriter, r *http.Request) {
		session.Save(r, w)
		response.Write(w, r)
	}

	return wrappedResponse
}

// example using decorator
func (h *IndexController) IndexWithDecorator(c *fireball.Context) (fireball.Response, error) {
	session := c.Meta["session"].(*sessions.Session)

	count, ok := session.Values["count"].(int)
	if !ok {
		count = 0
	}

	session.Values["count"] = count + 1
	body := fmt.Sprintf("You have visited this page %d time(s)", session.Values["count"])

	return fireball.NewResponse(200, []byte(body), nil), nil
}

// example rediect
func (h *IndexController) Redirect(c *fireball.Context) (fireball.Response, error) {
	var response fireball.ResponseFunc = func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://google.com", 301)
	}

	return response, nil
}
