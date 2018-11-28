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
