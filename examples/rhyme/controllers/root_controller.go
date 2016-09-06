package controllers

import (
	"github.com/zpatrick/fireball"
	"math/rand"
)

type RootController struct{}

func NewRootController() *RootController {
	return &RootController{}
}

func (h *RootController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/",
			Handlers: map[string]fireball.Handler{
				"GET": h.Index,
			},
		},
	}

	return routes
}

type Data struct {
	Lines []*Line
}

func (h *RootController) Index(c *fireball.Context) (fireball.Response, error) {
	lines := []*Line{}

	for len(lines) < 4 {
		song := Songs[rand.Intn(len(Songs))]
		if len(song.Lines) == 0 {
			continue
		}

		for {
			line := song.Lines[rand.Intn(len(song.Lines))]
			if len(line.Matches) == 0 {
				continue
			}

			match := line.Matches[rand.Intn(len(line.Matches))]
			lines = append(lines, line, match)
			break
		}
	}

	data := Data{
		Lines: lines,
	}

	return c.HTML(200, "index.html", data)
}
