package fireball

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Todo: Context should be interface that we can mock during testing
type Context struct {
	PathVariables map[string]string
	request       *http.Request
}

func (c *Context) HTML(status int, file string, data interface{}) (*HTMLResponse, error) {
	// todo: not hardcode "views" directory
	// todo: try subdirectories under views
	tmpl, err := c.parseTemplates()
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
			Headers:    map[string]string{"Content-Type": "text/html"},
		},
	}

	return response, nil
}

/*
func (c *Context) parseTemplates() (*template.Template, error) {
	paths := []string{}
	walkf := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			path = strings.Replace(path, "\\", "/", -1)
			paths = append(paths, path)
		}

		return nil
	}

	if err := filepath.Walk("views/", walkf); err != nil {
		return nil, err
	}

	tmpl, err := template.ParseFiles(paths...)
	if err != nil {
		return nil, err
	}

	return tmpl, nil

}
*/

func (c *Context) parseTemplates() (*template.Template, error) {
	root := template.New("root")

	walkf := func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			path = strings.Replace(path, "\\", "/", -1)

			current, err := template.ParseFiles(path)
			if err != nil {
				return err
			}

			for _, c := range current.Templates() {
				fmt.Println("Internal has templates: ", c.Name())

				/*
					var name string
					if i == 0 {
						filepath.Dir(path)
						// file template
						//name = strings.TrimPrefix(path, "views/")
						name = path
					} else {
						// custom template
						name = c.Name()
					}
				*/
				name := fmt.Sprintf("%s/%s", filepath.Dir(path), c.Name())

				if _, err := root.AddParseTree(name, c.Tree); err != nil {
					return err
				}
			}
		}

		/*

			glob := fmt.Sprintf("%s/%s", path, "*.html")
			t, err := template.ParseGlob(glob)
			if err != nil {
				return err
			}

			tmpl.AddParseTree(path, t.Tree)
		*/

		return nil
	}

	if err := filepath.Walk("views/", walkf); err != nil {
		return nil, err
	}

	for _, r := range root.Templates() {
		fmt.Println("Root has templates: ", r.Name())
	}

	return root, nil
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
			Headers:    map[string]string{"Content-Type": "application/json"},
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
