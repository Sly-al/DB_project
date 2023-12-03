package handlers

import (
	"DB_project/internal/storage"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func CustomerOrderPage(w http.ResponseWriter, r *http.Request) {
	data, err := storage.SelectCatalog()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("front", "pages", "customer", "make_order.html")
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

func CustomerCreateOrder(w http.ResponseWriter, r *http.Request) {
	catalog := r.FormValue("selectedCar")
	catalogId, err := strconv.Atoi(catalog)
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	err = storage.CreateNewOrder(catalogId, customerId, time.Now().AddDate(0, 0, 3))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

func CustomerAllOrders(w http.ResponseWriter, r *http.Request) {
	data, err := storage.SelectAllOrders(customerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("front", "pages", "customer", "orders.html")
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
