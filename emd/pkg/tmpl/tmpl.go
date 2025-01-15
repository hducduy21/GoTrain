package tmpl

import (
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig"
)

var Tmpl *template.Template

func ParseTemplates() error {
	log.Println("Parsing templates")
	tmp := template.New("").Funcs(sprig.FuncMap())
	err := filepath.Walk("frontend/templates", func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			tmplBytes, err := os.ReadFile(path)
			if err != nil {
				log.Println("err 1")
				return err
			}

			_, err = tmp.New(path).Funcs(sprig.FuncMap()).Parse(string(tmplBytes))
			if err != nil {
				log.Println("err 2")
				return err

			}
		}
		return err
	})
	if err != nil {
		return err
	}
	Tmpl = tmp
	return nil
}
