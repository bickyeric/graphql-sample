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
	_, err := database.Connection.Exec(
		`INSERT INTO item_category(id, name, outlet_id, store_id)
		VALUES (?, ?, ?, ?)`, ic.lastID()+1, ic.Name, ic.OutletID, ic.StoreID,
	)
	if err != nil {
		return ic, err
	}
	ic.ID = ic.lastID()
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

// GetByStoreID ...
func (ic ItemCategory) GetByStoreID() ([]ItemCategory, error) {
	var ics []ItemCategory
	rows, err := database.Connection.Query(
		`SELECT id, name, outlet_id, store_id
		FROM item_category WHERE store_id=?`, ic.StoreID,
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

func (ic *ItemCategory) lastID() int {
	row := database.Connection.QueryRow(
		`SELECT id
		FROM item_category
		WHERE store_id=? AND outlet_id=?
		ORDER BY id DESC
		LIMIT 1`, ic.StoreID, ic.OutletID,
	)
	id := 0
	row.Scan(&id)
	return id
}
