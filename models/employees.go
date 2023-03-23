package models

type Employee struct {
	Id    int
	Name  string
	Email string
}

const EmployeeSchema string = `CREATE TABLE employees (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
	email VARCHAR(50) NOT NULL,
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`
