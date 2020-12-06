package utils

import (
	"net/http"
	"text/template"

	"github.com/Hrtkk/helm-tutorial/models"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, p *models.Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}
