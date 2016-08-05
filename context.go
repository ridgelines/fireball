package fireball

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	HTMLHeaders = map[string]string{"Content-Type": "text/html"}
	JSONHeaders = map[string]string{"Content-Type": "application/json"}
)

// Todo: Context should be interface that we can mock during testing
type Context struct {
	PathVariables map[string]string
	Parser        TemplateParser
	request       *http.Request
}

func (c *Context) HTML(status int, file string, data interface{}) (*HTTPResponse, error) {
	tmpl, err := c.Parser.Parse()
	if err != nil {
		return nil, c.Error(500, err)
	}

	var buffer bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buffer, file, data); err != nil {
		return nil, c.Error(500, err)
	}

	response := NewHTTPResponse(status, buffer.Bytes(), HTMLHeaders)
	return response, nil
}

func (c *Context) JSON(status int, data interface{}) (*HTTPResponse, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response := NewHTTPResponse(status, bytes, JSONHeaders)
	return response, nil
}

func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) PathVar(key string) string {
	return c.PathVariables[key]
}

// todo
func (c *Context) QueryVar(key string) string {
	return ""
}

// todo
func (c *Context) FormVar(key string) string {
	return ""
}

func (c *Context) Error(status int, err error) *HTTPError {
	return NewHTTPError(status, err)
}

func (c *Context) Errorf(status int, format string, tokens ...interface{}) *HTTPError {
	return NewHTTPError(status, fmt.Errorf(format, tokens...))
}

func (c *Context) JSONError(status int, err error) *JSONError {
	return NewJSONError(status, err)
}

func (c *Context) JSONErrorf(status int, format string, tokens ...interface{}) *JSONError {
	return NewJSONError(status, fmt.Errorf(format, tokens...))
}
