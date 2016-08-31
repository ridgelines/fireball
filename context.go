package fireball

import (
	"bytes"
	"net/http"
)

var (
	HTMLHeaders = map[string]string{"Content-Type": "text/html"}
	JSONHeaders = map[string]string{"Content-Type": "application/json"}
)

type Context struct {
	PathVariables map[string]string
	Parser        TemplateParser
	request       *http.Request
}

func (c *Context) HTML(status int, file string, data interface{}) (*HTTPResponse, error) {
	tmpl, err := c.Parser.Parse()
	if err != nil {
		return nil, NewHTTPError(500, err, nil)
	}

	var buffer bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buffer, file, data); err != nil {
		return nil, NewHTTPError(500, err, nil)
	}

	response := NewHTTPResponse(status, buffer.Bytes(), HTMLHeaders)
	return response, nil
}

func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) PathVar(key string) string {
	return c.PathVariables[key]
}
