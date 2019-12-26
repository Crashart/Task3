package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Uniter interface {
	GetUnitList() ([]*Unit, error)
	GetUnitById(id int64)(*Unit, error)
	CreateUnit(details *Unit) *Unit
	UpdateUnit(id int64, details *Unit) *Unit
	DeleteUnit(id int64) error
}

type databaseU struct {
	*sql.DB
}

func Connect() *databaseU {
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
	return  &databaseU{db}
}

func (db *databaseU) GetUnitList() ([]*Unit, error) {
	rows, err := db.Query("SELECT * FROM units")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	units :=make([]*Unit, 0)
	for rows.Next() {
		var un Unit
		err := rows.Scan(&un.Id, &un.Name)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		units = append(units, &un)
	}
	if err = rows.Err(); err != nil{
			log.Println(err)
			return nil, err
	}
	return units, nil
}

func (db *databaseU) GetUnitById(id int64) (*Unit, error) {
	row, err := db.Query("SELECT * FROM units WHERE id=$1", id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer row.Close()
	var unit *Unit
	for row.Next() {
		err := row.Scan(&unit.Id, &unit.Name)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	if err = row.Err(); err != nil {
		return nil, err
	}
	return unit, nil
}

func (db *databaseU) CreateUnit(details *Unit) *Unit {
	_, err := db.Query("insert into units (name) values (&1)", details.Name)
	if err != nil {
		log.Println(err)
		return nil
	}
	return details
}

func (db *databaseU) UpdateUnit(id int64, details *Unit) *Unit {
	_, err := db.Query("UPDATE units SET name=$1 WHERE id=$2", details.Name, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return details
}

func (db *databaseU) DeleteUnit(id int64) error {
	_, err := db.Query("DELETE FROM units WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}