package views

import (
	"html/template"
	"log"
	"path/filepath"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = "tmpl"
)

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

// layoutFiles returns a slice of strings
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)

	if err != nil {
		log.Fatalf("-- could not Glob layouts: %v", err)
	}

	return files
}
