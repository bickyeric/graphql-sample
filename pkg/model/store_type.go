package model

import (
	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
)

// StoreType ...
type StoreType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Create ...
func (st StoreType) Create() (StoreType, error) {
	row, err := database.Connection.Exec(
		`INSERT INTO store_type(name) VALUES (?)`, st.Name,
	)
	if err != nil {
		return st, err
	}
	id, _ := row.LastInsertId()
	st.ID = int(id)
	return st, nil
}

// All ...
func (st StoreType) All() ([]StoreType, error) {
	var sts []StoreType
	rows, err := database.Connection.Query(
		`SELECT id, name
		FROM store_type`,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&st.ID, &st.Name)
		sts = append(sts, st)
	}

	return sts, nil
}

// GetByID ...
func (st StoreType) GetByID() (StoreType, error) {
	row := database.Connection.QueryRow(
		`SELECT id, name
		FROM store_type WHERE id=?`, st.ID,
	)
	err := row.Scan(&st.ID, &st.Name)
	if err != nil {
		return st, err
	}
	return st, nil
}

// Category ...
func (st *StoreType) Category() ([]StoreCategory, error) {
	return StoreCategory{TypeID: st.ID}.GetByType()
}

// Store ...
func (st *StoreType) Store() ([]Store, error) {
	return Store{TypeID: st.ID}.GetByType()
}
