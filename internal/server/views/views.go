package views

import (
	"embed"
	"html/template"
	"io"
)

//go:embed html/*.html
var content embed.FS

func Render(w io.Writer, templateName string, data interface{}) {
	templates := template.Must(template.ParseFS(content, "html/base.html", "html/"+templateName))
	templates.ExecuteTemplate(w, templateName, data)
}
