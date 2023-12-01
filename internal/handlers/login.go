package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	name string
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("front/pages/start.html")
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

func GetLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)
	if username == "admin" && password == "admin" {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		customer := user{username}
		http.Redirect(w, r, "/customer", http.StatusSeeOther)
		_ = customer
	}
}
