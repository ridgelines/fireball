package fireball

import (
	"bytes"
	"testing"
)

var golden = `<!DOCTYPE html>
<html>
<head>
    <title>Title</title>
</head>
<body>
    <h1>Header</h1>
    <h1>Hello, World!</h1>
    <h1>Footer</h1>
</body>
</html>
`

func TestGlobParse(t *testing.T) {
	parser := NewGlobParser("testing/", "*.html")

	tmpl, err := parser.Parse()
	if err != nil {
		t.Fatal(err)
	}

	var buffer bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buffer, "index.html", nil); err != nil {
		t.Fatal(err)
	}

	if v, want := buffer.String(), golden; v != want {
		t.Errorf("\nExpected: %#v \nReceived: %#v", want, v)
	}
}
