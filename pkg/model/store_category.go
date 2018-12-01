package model

import (
	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
)

// StoreCategory ...
type StoreCategory struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TypeID int    `json:"type_id"`
}

// Create ...
func (sc StoreCategory) Create() (StoreCategory, error) {
	_, err := database.Connection.Exec(
		`INSERT INTO store_category(id, name, type_id) VALUES (?, ?, ?)`, sc.ID, sc.Name, sc.TypeID,
	)
	if err != nil {
		return sc, err
	}
	return sc, nil
}

// All ...
func (sc StoreCategory) All() ([]StoreCategory, error) {
	var scs []StoreCategory
	rows, err := database.Connection.Query(
		`SELECT id, name, type_id
		FROM store_category`,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&sc.ID, &sc.Name, &sc.TypeID)
		scs = append(scs, sc)
	}

	return scs, nil
}

// GetByID ...
func (sc StoreCategory) GetByID() (StoreCategory, error) {
	row := database.Connection.QueryRow(
		`SELECT id, name, type_id
		FROM store_category WHERE id=?`, sc.ID,
	)
	err := row.Scan(&sc.ID, &sc.Name, &sc.TypeID)
	if err != nil {
		return sc, err
	}
	return sc, nil
}

// GetByType ...
func (sc StoreCategory) GetByType() ([]StoreCategory, error) {
	var scs []StoreCategory
	rows, err := database.Connection.Query(
		`SELECT id, name, type_id
		FROM store_category WHERE type_id=?`, sc.TypeID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&sc.ID, &sc.Name, &sc.TypeID)
		scs = append(scs, sc)
	}
	return scs, nil
}

// Type ...
func (sc StoreCategory) Type() (StoreType, error) {
	return StoreType{ID: sc.TypeID}.GetByID()
}
