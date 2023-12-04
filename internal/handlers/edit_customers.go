package handlers

import (
	"DB_project/internal/storage"
	"DB_project/internal/structers"
	"html/template"
	"log"
	"net/http"
)

func AdminEditCustomers(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("front/pages/admin/edit_customers.html")
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

func AdminAddClient(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("front/pages/admin/add_client.html")
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

func AdminInsertNewClient(w http.ResponseWriter, r *http.Request) {
	newclient := structers.Client{}
	newclient.Login = r.FormValue("customerLogin")
	newclient.Password = r.FormValue("customerPassword")
	newclient.Surname = r.FormValue("customerSurname")
	newclient.Name = r.FormValue("customerName")
	regular := r.FormValue("statusRegular")
	vip := r.FormValue("statusVIP")
	if regular == "" {
		newclient.Status = vip
	} else {
		newclient.Status = regular
	}
	err := storage.InsertNewClient(newclient)
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

func AdminRemoveClient(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("front/pages/admin/remove_client.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Fatalf("StartPage: %s", err.Error())
	}
	logins, err := storage.SelectAllClientLogins()
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}

	data := struct {
		Logins []string
	}{
		logins,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalf("StartPage: %s", err.Error())
	}
}

func AdminDeleteClient(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	err := storage.DeleteClient(login)
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

func AdminVIPPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("front/pages/admin/give_vip.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		log.Fatalf("StartPage: %s", err.Error())
	}
	logins, err := storage.SelectAllRegularClientLogins()
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}

	data := struct {
		Logins []string
	}{
		logins,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalf("StartPage: %s", err.Error())
	}
}

func AdminSetVIPStatus(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	err := storage.UpdateStatus(login)
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}
