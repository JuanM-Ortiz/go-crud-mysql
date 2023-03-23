package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/JuanM-Ortiz/go-crud-mysql/db"
	"github.com/JuanM-Ortiz/go-crud-mysql/models"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func Index(rw http.ResponseWriter, r *http.Request) {
	connectionOn := db.Connect()

	records, err := connectionOn.Query("SELECT id, name, email FROM employees")

	if err != nil {
		panic(err.Error())
	}

	employee := models.Employee{}
	empArray := []models.Employee{}

	for records.Next() {
		var id int
		var name, email string
		err = records.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}

		employee.Id = id
		employee.Name = name
		employee.Email = email

		empArray = append(empArray, employee)
	}

	templates.ExecuteTemplate(rw, "index", empArray)

}

func Create(rw http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(rw, "create", nil)
}

func Insert(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		name := r.FormValue("name")
		email := r.FormValue("email")

		connectionOn := db.Connect()

		insertRecords, err := connectionOn.Prepare("INSERT INTO employees(name, email) VALUES(?,?)")

		if err != nil {
			panic(err)
		}

		insertRecords.Exec(name, email)

		http.Redirect(rw, r, "/", http.StatusMovedPermanently)

	}
}

func Edit(rw http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")

	connectionOn := db.Connect()

	record, err := connectionOn.Query("SELECT id, name, email FROM employees WHERE id=?", idEmployee)

	if err != nil {
		panic(err.Error())
	}

	employee := models.Employee{}

	for record.Next() {
		var id int
		var name, email string
		err = record.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}

		employee.Id = id
		employee.Name = name
		employee.Email = email
	}

	templates.ExecuteTemplate(rw, "edit", employee)

}

func Update(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")

		connectionOn := db.Connect()

		updateRecords, err := connectionOn.Prepare("UPDATE employees SET name=?, email=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		updateRecords.Exec(name, email, id)

		http.Redirect(rw, r, "/", http.StatusMovedPermanently)

	}
}

func Delete(rw http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")
	fmt.Println(idEmployee)

	connectionOn := db.Connect()

	deleteRecords, err := connectionOn.Prepare("DELETE FROM employees WHERE id=?")

	if err != nil {
		panic(err)
	}

	deleteRecords.Exec(idEmployee)

	http.Redirect(rw, r, "/", http.StatusMovedPermanently)

}
