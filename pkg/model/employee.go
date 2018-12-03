package model

import (
	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
)

// Employee ...
type Employee struct {
	ID          int    `json:"id"`
	OutletID    int    `json:"outlet_id"`
	StoreID     int    `json:"store_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Confirmed   bool   `json:"confirmed"`
	Active      bool   `json:"active"`
}

// Create ...
func (e Employee) Create() (Employee, error) {
	row, err := database.Connection.Exec(
		`INSERT INTO employee(outlet_id, store_id, first_name, last_name, phone_number, email, password, confirmed, active)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`, e.OutletID, e.StoreID, e.FirstName, e.LastName, e.PhoneNumber, e.Email, e.Password, e.Confirmed, e.Active,
	)
	if err != nil {
		return e, err
	}
	id, _ := row.LastInsertId()
	e.ID = int(id)
	return e, nil
}

// All ...
func (e Employee) All() ([]Employee, error) {
	var employees []Employee
	rows, err := database.Connection.Query(
		`SELECT id, outlet_id, store_id, first_name, last_name, phone_number, email, password, confirmed, active
		FROM employee`,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&e.ID, &e.OutletID, &e.StoreID, &e.FirstName, &e.LastName, &e.PhoneNumber, &e.Email, &e.Password, &e.Confirmed, &e.Active)
		employees = append(employees, e)
	}

	return employees, nil
}

// GetByOutletID ...
func (e Employee) GetByOutletID() ([]Employee, error) {
	var employees []Employee
	rows, err := database.Connection.Query(
		`SELECT id, outlet_id, store_id, first_name, last_name, phone_number, email, password, confirmed, active
		FROM employee
		WHERE outlet_id=? AND store_id=?`, e.OutletID, e.StoreID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&e.ID, &e.OutletID, &e.StoreID, &e.FirstName, &e.LastName, &e.PhoneNumber, &e.Email, &e.Password, &e.Confirmed, &e.Active)
		employees = append(employees, e)
	}

	return employees, nil
}

// GetByStoreID ...
func (e Employee) GetByStoreID() ([]Employee, error) {
	var employees []Employee
	rows, err := database.Connection.Query(
		`SELECT id, outlet_id, store_id, first_name, last_name, phone_number, email, password, confirmed, active
		FROM employee
		WHERE store_id=?`, e.StoreID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&e.ID, &e.OutletID, &e.StoreID, &e.FirstName, &e.LastName, &e.PhoneNumber, &e.Email, &e.Password, &e.Confirmed, &e.Active)
		employees = append(employees, e)
	}

	return employees, nil
}

// Outlet ...
func (e Employee) Outlet() (Outlet, error) {
	return Outlet{ID: e.OutletID, StoreID: e.StoreID}.GetByID()
}

// Store ...
func (e Employee) Store() (Store, error) {
	return Store{ID: e.StoreID}.GetByID()
}
