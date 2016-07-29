package fireball

import (
	"net/http"
	"strings"
)

type Handler func(c *Context) (interface{}, error)

type Route struct {
	Path string
	Get  Handler
	Post Handler
}

// todo: return Handler+Variables
func (r *Route) Match(req *http.Request) *Match {
	variables, ok := r.matchPath(req.URL.Path)
	if !ok {
		return nil
	}

	var handler Handler
	switch req.Method {
	case "GET":
		handler = r.Get
	case "POST":
		handler = r.Post
	}

	if handler == nil {
		return nil
	}

	match := &Match{
		Route:     r,
		Handler:   handler,
		Variables: variables,
	}

	return match
}

func (r *Route) matchPath(url string) (map[string]string, bool) {
	pathSections := strings.Split(r.Path, "/")
	urlSections := strings.Split(url, "/")

	if len(pathSections) != len(urlSections) {
		return nil, false
	}

	// todo: special case on trailing /
	variables := map[string]string{}
	for i, pathSection := range pathSections {
		urlSection := urlSections[i]

		if strings.HasPrefix(pathSection, "{") && strings.HasSuffix(pathSection, "}") {
			key := pathSection[1 : len(pathSection)-1]
			variables[key] = urlSection
		} else if pathSection != urlSection {
			return nil, false
		}
	}

	return variables, true

}
