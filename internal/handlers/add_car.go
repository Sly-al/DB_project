package handlers

import (
	"DB_project/internal/storage"
	"DB_project/internal/structers"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func AdminAddCar(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("front/pages/admin/add_car.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Fatalf("StartPage: %s", err.Error())
	}
	brands, err := storage.SelectAllBrands()
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Fatalf("StartPage: %s", err.Error())
	}

	equipmets, err := storage.SelectAllEquipment()
	var Engine, Color, Transmission, Body []string
	for _, now := range equipmets {
		Engine = append(Engine, now.Engine)
		Color = append(Color, now.Color)
		Transmission = append(Transmission, now.Transmission)
		Body = append(Body, now.Body)
	}
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Fatalf("StartPage: %s", err.Error())
	}
	data := struct {
		Brands       []string
		Engine       []string
		Color        []string
		Transmission []string
		Body         []string
	}{
		brands,
		Engine,
		Color,
		Transmission,
		Body,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalf("StartPage: %s", err.Error())
	}
}

func AddNewCar(w http.ResponseWriter, r *http.Request) {
	var eq structers.Equipment
	var newCar bool
	carName := r.FormValue("carName")
	isnew := r.FormValue("newCar")
	if isnew == "yes" {
		newCar = true
	} else {
		newCar = false
	}
	brand := r.FormValue("carBrand")
	eq.Body = r.FormValue("carBody")
	eq.Color = r.FormValue("carColor")
	eq.Transmission = r.FormValue("carTransmission")
	eq.Engine = r.FormValue("carEngine")

	eqId, err := storage.InsertToEq(eq)
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}

	brandId, err := storage.SelectIdSup(brand)
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	nm := structers.Machine{
		0,
		carName,
		newCar,
		brandId,
		eqId,
	}
	fmt.Println(nm)
	err = storage.InsertNewMachine(nm)
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}
