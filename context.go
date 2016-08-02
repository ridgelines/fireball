package fireball

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Todo: Context should be interface that we can mock during testing
type Context struct {
	PathVariables map[string]string
	request       *http.Request
}

func (c *Context) HTML(status int, file string, data interface{}) (*HTMLResponse, error) {
	tmpl, err := template.ParseGlob("*.html")
	if err != nil {
		return nil, c.Error(500, err)
	}

	var buffer bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buffer, file, data); err != nil {
		return nil, c.Error(500, err)
	}

	response := &HTMLResponse{
		&HTTPResponse{
			Data:       buffer.Bytes(),
			StatusCode: status,
		},
	}

	return response, nil
}

func (c *Context) JSON(status int, data interface{}) (*JSONResponse, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response := &JSONResponse{
		&HTTPResponse{
			Data:       bytes,
			StatusCode: status,
		},
	}

	return response, nil
}

func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) UnmarshalBodyJSON(obj interface{}) error {
	decoder := json.NewDecoder(c.request.Body)
	return decoder.Decode(&obj)
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
	return &HTTPError{
		StatusCode: status,
		Err:        err,
	}
}

func (c *Context) Errorf(status int, format string, tokens ...interface{}) *HTTPError {
	return c.Error(status, fmt.Errorf(format, tokens...))
}

// todo: HTMLError

func (c *Context) JSONError(status int, err error) *JSONError {
	return &JSONError{
		&HTTPError{
			StatusCode: status,
			Err:        err,
		},
	}
}

func (c *Context) JSONErrorf(status int, format string, tokens ...interface{}) *JSONError {
	return c.JSONError(status, fmt.Errorf(format, tokens...))
}
