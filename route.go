package fireball

type Handler func(c *Context) (interface{}, error)

type Route struct {
	Path     string
	Handlers map[string]Handler
}
