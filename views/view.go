package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)
var (
	LayoutDir string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return  &View{
		Template: t,
		Layout: layout,
	}
}

type View struct {
	Template *template.Template
	Layout string
}

//Render is used to render the view with the predefined layout.
func (v *View) Render(w http.ResponseWriter, data interface{}) error{
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

//Layout files return a slice of string representing
// the layout file using in our apps
func layoutFiles() []string{
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err !=nil {
		panic(err)
	}
	return files
}
