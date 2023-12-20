package handlers

import (
	"DB_project/internal/storage"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func AdminRemoveCar(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("front/pages/admin/remove_car.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Fatalf("StartPage: %s", err.Error())
	}
	data, err := storage.SelectCatalog()
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Fatalf("StartPage: %s", err.Error())
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalf("StartPage: %s", err.Error())
	}
}

func AdminDeleteCar(w http.ResponseWriter, r *http.Request) {
	catalogId, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	err = storage.DeleteCatalog(catalogId)
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}
