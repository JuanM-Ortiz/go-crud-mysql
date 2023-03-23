package main

import (
	"log"
	"net/http"

	"github.com/JuanM-Ortiz/go-crud-mysql/db"
	"github.com/JuanM-Ortiz/go-crud-mysql/models"
	"github.com/JuanM-Ortiz/go-crud-mysql/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.Index)
	router.HandleFunc("/create", routes.Create)
	router.HandleFunc("/insert", routes.Insert)
	router.HandleFunc("/delete", routes.Delete)
	router.HandleFunc("/edit", routes.Edit)
	router.HandleFunc("/update", routes.Update)

	db.Connect()
	db.CreateTable(models.EmployeeSchema, "employees")

	log.Fatal(http.ListenAndServe(":8080", router))
}
