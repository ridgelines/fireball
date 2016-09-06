package fireball

type Handler func(c *Context) (Response, error)

type Route struct {
	Path     string
	Handlers map[string]Handler
}
