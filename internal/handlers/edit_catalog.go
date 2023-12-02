package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func AdminEditCatalog(w http.ResponseWriter, t *http.Request) {
	tmpl, err := template.ParseFiles("front/pages/admin/edit_catalog.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Fatalf("StartPage: %s", err.Error())
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalf("StartPage: %s", err.Error())
	}
}
