package model

import "github.com/bickyeric/graphql-sample/pkg/wrapper/database"

// StoreType ...
type StoreType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Create ...
func (st StoreType) Create() (StoreType, error) {

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
