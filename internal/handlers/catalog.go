package handlers

import (
	"DB_project/internal/storage"
	"html/template"
	"net/http"
	"path/filepath"
)

func AdminCatalog(w http.ResponseWriter, r *http.Request) {
	data, err := storage.SelectCatalog()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("front", "pages", "admin", "catalog.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func CustomerCatalog(w http.ResponseWriter, r *http.Request) {
	data, err := storage.SelectCatalog()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("front", "pages", "customer", "catalog.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
