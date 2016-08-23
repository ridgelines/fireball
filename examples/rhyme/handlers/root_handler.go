package handlers

import (
	"github.com/zpatrick/fireball"
	"math/rand"
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

type Context struct {
	Lines []*Line
}

func (h *RootHandler) Index(c *fireball.Context) (interface{}, error) {
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

	context := Context{
		Lines: lines,
	}

	return c.HTML(200, "index.html", context)
}
