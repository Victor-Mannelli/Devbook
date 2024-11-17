package utils

import (
	"devbook/src/responses"
	"net/http"
	"text/template"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*html"))
	templates = template.Must(templates.ParseGlob("views/templates/*html"))
}

func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	err := templates.ExecuteTemplate(w, template, data)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
}
