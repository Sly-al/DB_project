package handlers

import (
	"DB_project/internal/storage"
	"html/template"
	"log"
	"net/http"
)

var customerId int

//
//func LoginPage(w http.ResponseWriter, r *http.Request) {
//	tmpl, err := template.ParseFiles("front/pages/start.html")
//	if err != nil {
//		http.Error(w, err.Error(), 400)
//		log.Fatalf("StartPage: %s", err.Error())
//	}
//	err = tmpl.Execute(w, nil)
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		log.Fatalf("StartPage: %s", err.Error())
//	}
//}
//
//func GetLogin(w http.ResponseWriter, r *http.Request) {
//	login := r.FormValue("username")
//	passwordRec := r.FormValue("password")
//	if login == "admin" && passwordRec == "admin" {
//		http.Redirect(w, r, "/admin", http.StatusSeeOther)
//	} else {
//		passwordReal, err := storage.GetPassword(login)
//		if err != nil {
//			http.Error(w, err.Error(), 400)
//		}
//		if passwordRec == passwordReal {
//			http.Redirect(w, r, "/customer", http.StatusSeeOther)
//		} else {
//			http.Error(w, "No such customer or client", 400)
//		}
//	}
//}

type LoginPageData struct {
	Error string
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	data := LoginPageData{}
	tmpl, err := template.ParseFiles("front/pages/start.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("LoginPage:", err)
		return
	}

	if r.Method == http.MethodPost {
		login := r.FormValue("username")
		passwordRec := r.FormValue("password")

		if login == "admin" && passwordRec == "admin" {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		passwordReal, err := storage.GetPassword(login)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("LoginPage:", err)
			return
		}

		if passwordRec == passwordReal {
			http.Redirect(w, r, "/customer", http.StatusSeeOther)
			return
		}

		data.Error = "Неправильный логин или пароль"
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("LoginPage:", err)
		return
	}
}
