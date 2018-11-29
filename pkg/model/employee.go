package model

import (
	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
)

// Employee ...
type Employee struct {
	ID          int
	OutletID    int
	StoreID     int
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Password    string
	Confirmed   bool
	Active      bool
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

// GetByOutletID ...
func (e Employee) GetByOutletID() ([]Employee, error) {
	var employees []Employee
	rows, err := database.Connection.Query(
		`SELECT outlet_id, store_id, first_name, last_name, phone_number, email, password, confirmed, active
		FROM employee
		WHERE outlet_id=? AND store_id=?`, e.OutletID, e.StoreID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&e.OutletID, &e.StoreID, &e.FirstName, &e.LastName, &e.PhoneNumber, &e.Email, &e.Password, &e.Confirmed, &e.Active)
		employees = append(employees, e)
	}

	return employees, nil
}
