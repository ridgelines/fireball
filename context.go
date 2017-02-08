package fireball

import (
	"net/http"
)

var (
	HTMLHeaders = map[string]string{"Content-Type": "text/html"}
	JSONHeaders = map[string]string{"Content-Type": "application/json"}
	TextHeaders = map[string]string{"Content-Type": "text/plain"}
	CORSHeaders = map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "*",
		"Access-Control-Allow-Methods": "GET, POST, PUT, PATCH, DELETE, COPY, HEAD, OPTIONS, LINK, UNLINK, CONNECT, TRACE, PURGE",
	}
)

// Context is passed into Handlers
// It contains fields and helper functions related to the request
type Context struct {
	// PathVariables are the URL-related variables returned by the Router
	PathVariables map[string]string
	// Meta can be used to pass information along Decorators
	Meta map[string]interface{}
	// Parser is used to render html templates
	Parser TemplateParser
	// Request is the originating *http.Request
	Request *http.Request
}

// Context.HTML calls HTML with the Context's template parser
func (c *Context) HTML(status int, templateName string, data interface{}) (*HTTPResponse, error) {
	return HTML(c.Parser, status, templateName, data)
}
