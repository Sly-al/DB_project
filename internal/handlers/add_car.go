package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func AdminAddCar(w http.ResponseWriter, t *http.Request) {
	tmpl, err := template.ParseFiles("front/pages/admin/add_car.html")
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
