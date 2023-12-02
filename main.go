package main

import (
	"DB_project/internal/config"
	"DB_project/internal/handlers"
	"DB_project/internal/storage"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	err := storage.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	//err = storage.CreateNewOrder(3, time.Now().AddDate(0, 0, 7))
	//if err != nil {
	//	fmt.Println(err)
	//}

	router := chi.NewRouter()
	router.Handle("/front/*", http.StripPrefix("/front/", http.FileServer(http.Dir("front"))))
	router.Get("/", handlers.LoginPage)
	router.Post("/", handlers.LoginPage)

	router.Get("/admin", handlers.AdminStartPage)
	router.Get("/admin/catalog", handlers.AdminCatalog)
	router.Get("/admin/edit_catalog", handlers.AdminEditCatalog)
	router.Get("/admin/add_car", handlers.AdminAddCar)
	//router.Get("/admin/remove_car", handlers.AdminRemoveCar)
	router.Get("/admin/edit_customers", handlers.AdminEditCustomers)
	router.Get("/admin/add_client", handlers.AdminAddClient)
	router.Post("/admin/add_client", handlers.AdminInsertNewClient)

	router.Get("/customer", handlers.CustomerStartPage)
	router.Get("/customer/catalog", handlers.CustomerCatalog)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
