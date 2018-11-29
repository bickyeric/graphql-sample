package model

import "github.com/bickyeric/graphql-sample/pkg/wrapper/database"

// Variant ...
type Variant struct {
	ID         int
	ItemID     int
	Price      int
	SKU        string
	Stock      int
	TrackStock bool
	Alert      bool
	AlertAt    int
	Cost       int
	TrackCOGS  bool
}

// Create ...
func (v Variant) Create() (Variant, error) {
	row, err := database.Connection.Exec(
		`INSERT INTO variant(item_id, price, sku, stock, track_stock, alert, alert_at, cost, track_cogs)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`, v.ItemID, v.Price, v.SKU, v.Stock, v.TrackCOGS, v.Alert, v.AlertAt, v.Cost, v.TrackCOGS,
	)
	if err != nil {
		return v, err
	}
	id, _ := row.LastInsertId()
	v.ID = int(id)
	return v, nil
}
