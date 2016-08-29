package fireball

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type TemplateParser interface {
	Parse() (*template.Template, error)
}

type GlobParser struct {
	Root string
	Glob string
}

func (p *GlobParser) Parse() (*template.Template, error) {
	root := template.New("root")

	walkf := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path = filepath.Join(path, p.Glob)
			current, err := template.ParseGlob(path)
			if err != nil {
				return err
			}

			for _, t := range current.Templates() {
				name := p.generateTemplateName(path, t)

				if _, err := root.AddParseTree(name, t.Tree); err != nil {
					return err
				}
			}
		}

		return nil
	}

	if err := filepath.Walk(p.Root, walkf); err != nil {
		return nil, err
	}

	return root, nil
}

func (p *GlobParser) generateTemplateName(path string, t *template.Template) string {
	path = strings.Replace(filepath.Dir(path), "\\", "/", -1)
	path = fmt.Sprintf("%s/%s", path, t.Name())
	name := strings.TrimPrefix(path, p.Root)
	return name
}
