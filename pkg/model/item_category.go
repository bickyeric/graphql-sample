package model

import "github.com/bickyeric/graphql-sample/pkg/wrapper/database"

// ItemCategory ...
type ItemCategory struct {
	ID       int
	Name     string
	OutletID int
	StoreID  int
}

// Create ...
func (ic ItemCategory) Create() (ItemCategory, error) {
	row, err := database.Connection.Exec(
		`INSERT INTO item_category(name, outlet_id, store_id)
		VALUES (?, ?, ?)`, ic.Name, ic.OutletID, ic.StoreID,
	)
	if err != nil {
		return ic, err
	}
	id, _ := row.LastInsertId()
	ic.ID = int(id)
	return ic, nil
}

// GetByOutletID ...
func (ic ItemCategory) GetByOutletID() ([]ItemCategory, error) {
	var ics []ItemCategory
	rows, err := database.Connection.Query(
		`SELECT id, name, outlet_id, store_id
		FROM item_category WHERE outlet_id=? AND store_id=?`, ic.OutletID, ic.StoreID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&ic.ID, &ic.Name, &ic.OutletID, &ic.StoreID)
		ics = append(ics, ic)
	}
	return ics, nil
}
