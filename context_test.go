package fireball

import (
	"html/template"
	"testing"
)

func TestHTML(t *testing.T) {
	parser := TemplateParserFunc(func() (*template.Template, error) {
		tmpl, err := template.New("template_name").Parse("<h1>{{ . }}</h1>")
		if err != nil {
			t.Fatal(err)
		}

		return tmpl, nil
	})

	context := &Context{Parser: parser}
	response, err := context.HTML(200, "template_name", "some data")
	if err != nil {
		t.Fatal(err)
	}

	if v, want := string(response.Body), "<h1>some data</h1>"; v != want {
		t.Errorf("\nExpected: %#v \nReceived: %#v", want, v)
	}
}
