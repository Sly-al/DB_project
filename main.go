package main

import (
	"DB_project/internal/config"
	"DB_project/internal/storage"
	"fmt"
	"log"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	err := storage.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	err = storage.UpdatePriceSale(1, 200, 2)
	if err != nil {
		fmt.Println(err)
	}

	//router := chi.NewRouter()
	//router.Handle("/front/*", http.StripPrefix("/front/", http.FileServer(http.Dir("front"))))
	//router.Get("/", handlers.LoginPage)
	//router.Post("/", handlers.GetLogin)
	//
	//router.Get("/admin", handlers.AdminStartPage)
	//router.Get("/admin/catalog", handlers.AdminCatalog)
	//
	//router.Get("/customer", handlers.CustomerStartPage)
	//router.Get("/customer/catalog", handlers.CustomerCatalog)
	//
	//if err := http.ListenAndServe(":8080", router); err != nil {
	//	log.Fatal(err)
	//}
}
