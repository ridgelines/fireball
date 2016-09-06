package fireball

// Handler performs the business logic on a request
type Handler func(c *Context) (Response, error)

// Routes are used to map a request to a RouteMatch
type Route struct {
	// Path is used to determine if a request's URL matches this Route
	Path string
	// Handlers map common HTTP methods to different Handlers
	Handlers map[string]Handler
}

// RouteMatch objects are returned by the router when a request is successfully matched
type RouteMatch struct {
	Handler       Handler
	PathVariables map[string]string
}
