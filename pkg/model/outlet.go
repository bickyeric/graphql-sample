package model

import "github.com/bickyeric/garut/pkg/wrapper/database"

// Outlet ...
type Outlet struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Status      bool   `json:"status"`
	Zip         string `json:"zip"`
	StoreID     int    `json:"store_id"`
}

// All ...
func (o Outlet) All(first, offset int) ([]Outlet, error) {
	if offset < 1 {
		offset = 20
	}
	var outlets []Outlet
	rows, err := database.Connection.Query(
		`SELECT id, name, address, phone_number, city, state, status, zip, store_id
		FROM outlet LIMIT ?,?`, first, offset,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&o.ID, &o.Name, &o.Address, &o.PhoneNumber, &o.City, &o.State, &o.Status, &o.Zip, &o.StoreID)
		outlets = append(outlets, o)
	}

	return outlets, nil
}

// GetByStoreID ...
func (o Outlet) GetByStoreID() ([]Outlet, error) {
	var outlets []Outlet
	rows, err := database.Connection.Query(
		`SELECT id, name, address, phone_number, city, state, status, zip, store_id
		FROM outlet WHERE store_id=?`, o.StoreID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&o.ID, &o.Name, &o.Address, &o.PhoneNumber, &o.City, &o.State, &o.Status, &o.Zip, &o.StoreID)
		outlets = append(outlets, o)
	}

	return outlets, nil
}

// Store ...
func (o *Outlet) Store() (Store, error) {
	return Store{ID: o.StoreID}.GetByID()
}
