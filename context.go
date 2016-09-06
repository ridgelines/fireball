package fireball

import (
	"bytes"
	"net/http"
)

var (
	HTMLHeaders = map[string]string{"Content-Type": "text/html"}
	JSONHeaders = map[string]string{"Content-Type": "application/json"}
	TextHeaders = map[string]string{"Content-Type": "text/plain"}
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
	// Writer is the originating http.ResponseWriter
	Writer http.ResponseWriter
}

// HTML is a helper function that returns a response generated from the given templateName and data
func (c *Context) HTML(status int, templateName string, data interface{}) (*HTTPResponse, error) {
	tmpl, err := c.Parser.Parse()
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buffer, templateName, data); err != nil {
		return nil, err
	}

	response := NewResponse(status, buffer.Bytes(), HTMLHeaders)
	return response, nil
}
