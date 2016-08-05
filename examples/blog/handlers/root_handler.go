package handlers

import (
	"github.com/zpatrick/fireball"
)

type RootHandler struct{}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (h *RootHandler) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		&fireball.Route{
			Path: "/",
			Get:  h.Index,
		},
	}

	return routes
}

func (h *RootHandler) Index(c *fireball.Context) (interface{}, error) {
	return c.HTML(200, "index.html", "data")
}
