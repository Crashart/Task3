package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Employer interface {
	GetEmployeeList(unit_id int64) ([]*Employee, error)
	GetEmployeeById(unit_id, empl_id int64) (*Employee, error)
	CreateEmployee(unit_id int64, details *Employee) *Employee
	UpdateEmployee(unit_id, empl_id int64, details *Employee) *Employee
	DeleteEmployee(unit_id, empl_id int64) error
}

type databaseE struct {
	*sql.DB
}

func ConnectToDb() *databaseE {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=resources_department sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Println(err)
		return nil
	}
	if err = db.Ping(); err != nil {
		log.Println(err)
		return nil
	}
	return &databaseE{db}
}

func (db *databaseE) GetEmployeeList(unit_id int64) ([]*Employee, error) {
	rows, err := db.Query("SELECT * FROM employees WHERE unit_id=$1", unit_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	employees := make([]*Employee, 0)
	for rows.Next() {
		var em Employee
		err := rows.Scan(&em.Id, &em.Name)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		employees = append(employees, &em)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return employees, nil
}

func (db *databaseE) GetEmployeeById(unit_id, empl_id int64) (*Employee, error) {
	row, err := db.Query("SELECT * FROM employees WHERE unit_id=$1 and id=$2", unit_id, empl_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer row.Close()
	var employee *Employee
	for row.Next() {
		err := row.Scan(&employee.Id, &employee.Name)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	if err = row.Err(); err != nil {
		return nil, err
	}
	return employee, nil
}

func (db *databaseE) CreateEmployee(unit_id int64, details *Employee) *Employee {
	_, err := db.Query("insert into employees (name, age, unit_id) values (&1, &2, &3)", details.Name, details.Age, unit_id)
	if err != nil {
		log.Println(err)
		return nil
	}
	return details
}

func (db *databaseE) UpdateEmployee(unit_id, empl_id int64, details *Employee) *Employee {
	_, err := db.Query("UPDATE employees SET name=$1, age=$2, unit_id=$3 WHERE unit_id=$4 and id=$5", details.Name, details.Age, details.Unit_id, unit_id, empl_id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return details
}

func (db *databaseE) DeleteEmployee(unit_id, empl_id int64) error {
	_, err := db.Query("DELETE FROM employees WHERE id=$1 and unit_id=$2", empl_id, unit_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
