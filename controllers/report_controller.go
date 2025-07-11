package controllers

import (
	"html/template"
	"net/http"
)

func ReportIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/report/index.html"))
	tmpl.Execute(w, nil)
}
