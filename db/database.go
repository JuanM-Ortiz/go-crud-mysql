package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "root:123456@tcp(localhost:3306)/go_crud_empleados"

var db *sql.DB

func Connect() (connection *sql.DB) {
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexion exitosa")
	db = connection

	return db
}

// Verificar conexion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// Verificar si existe una tabla
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)

	if err != nil {
		fmt.Println("Error:", err)
	}

	return rows.Next()
}

// Crea una tabla
func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := db.Exec(schema)

		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

// Reiniciar los registros de la tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	db.Exec(sql)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)

	if err != nil {
		fmt.Println("Error:", err)
	}

	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)

	if err != nil {
		fmt.Println("Error:", err)
	}

	return rows, err
}
