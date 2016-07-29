package fireball

import (
	"bytes"
	"encoding/json"
	"html/template"
)

type Context struct {
	PathVariables map[string]string
}

func (c *Context) HTML(status int, file string, data interface{}) (interface{}, error) {
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

// todo: HTMLError

func (c *Context) JSONError(status int, err error) *JSONError {
	return &JSONError{
		&HTTPError{
			StatusCode: status,
			Err:        err,
		},
	}
}
