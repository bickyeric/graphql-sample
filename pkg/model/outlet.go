package model

import (
	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
)

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

// Create ...
func (o Outlet) Create() (Outlet, error) {
	_, err := database.Connection.Exec(
		`INSERT INTO outlet(id, name, address, phone_number, city, state, status, zip, store_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`, o.lastID()+1, o.Name, o.Address, o.PhoneNumber, o.City, o.State, o.Status, o.Zip, o.StoreID,
	)
	if err != nil {
		return o, err
	}
	o.ID = o.lastID()
	return o, nil
}

// All ...
func (o Outlet) All() ([]Outlet, error) {
	var outlets []Outlet
	rows, err := database.Connection.Query(
		`SELECT id, name, address, phone_number, city, state, status, zip, store_id
		FROM outlet`,
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

// GetByID ...
func (o Outlet) GetByID() (Outlet, error) {
	err := database.Connection.QueryRow(
		`SELECT id, name, address, phone_number, city, state, status, zip, store_id
		FROM outlet
		WHERE id=? AND store_id=?`, o.ID, o.StoreID,
	).Scan(&o.ID, &o.Name, &o.Address, &o.PhoneNumber, &o.City, &o.State, &o.Status, &o.Zip, &o.StoreID)
	if err != nil {
		return o, err
	}

	return o, nil
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

// ItemCategory ...
func (o *Outlet) ItemCategory() ([]ItemCategory, error) {
	return ItemCategory{StoreID: o.StoreID, OutletID: o.ID}.GetByOutletID()
}

// Employee ...
func (o *Outlet) Employee() ([]Employee, error) {
	return Employee{StoreID: o.StoreID, OutletID: o.ID}.GetByOutletID()
}

func (o *Outlet) lastID() int {
	row := database.Connection.QueryRow(
		`SELECT id
		FROM outlet
		WHERE store_id=?
		ORDER BY id DESC
		LIMIT 1`, o.StoreID,
	)
	id := 0
	row.Scan(&id)
	return id
}
