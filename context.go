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
	Request       *http.Request
	Writer        http.ResponseWriter
	Meta          map[string]interface{}
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
