package controllers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/zpatrick/fireball"
)

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
		&fireball.Route{
			Path: "/",
			Handlers: map[string]fireball.Handler{
				"GET": i.Index,
			},
		},
	}

	return routes
}

func (h *IndexController) Index(c *fireball.Context) (interface{}, error) {
	session := c.Meta["session"].(*sessions.Session)

	count, ok := session.Values["count"].(int)
	if !ok {
		count = 0
	}

	count += 1
	session.Values["count"] = count
	return fmt.Sprintf("You have visited this page %d times", count), nil
}
