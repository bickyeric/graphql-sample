package model

import (
	"database/sql"

	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
	"github.com/graphql-go/graphql"
)

// Item ...
type Item struct {
	ID          int    `json:"id"`
	OutletID    int    `json:"outlet_id"`
	StoreID     int    `json:"store_id"`
	Name        string `json:"name"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description"`
	ModifierID  int    `json:"modifier_id"`
	Image       string `json:"image"`

	categoryID sql.NullInt64
	modifierID sql.NullInt64
	image      sql.NullString
}

// Create ...
func (i Item) Create() (Item, error) {
	row, err := database.Connection.Exec(
		`INSERT INTO item(outlet_id, store_id, name, category_id, description, modifier_id, image)
		VALUES (?, ?, ?, ?, ?, ?, ?)`, i.OutletID, i.StoreID, i.Name, NewNullInt64(i.CategoryID), i.Description, NewNullInt64(i.ModifierID), NewNullString(i.Image),
	)
	if err != nil {
		return i, err
	}
	id, _ := row.LastInsertId()
	i.ID = int(id)
	return i, nil
}

// ItemType ...
var ItemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Customer",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"image": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

// All ...
func (i Item) All() ([]Item, error) {
	var items []Item
	rows, err := database.Connection.Query(
		`SELECT id, outlet_id, store_id, name, category_id, description, modifier_id, image
		FROM item`,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i Item
		err := rows.Scan(&i.ID, &i.OutletID, &i.StoreID, &i.Name, &i.categoryID, &i.Description, &i.modifierID, &i.image)
		if err != nil {
			return nil, err
		}
		if i.categoryID.Valid {
			i.CategoryID = int(i.categoryID.Int64)
		}
		if i.modifierID.Valid {
			i.ModifierID = int(i.modifierID.Int64)
		}
		if i.image.Valid {
			i.Image = i.image.String
		}
		items = append(items, i)
	}

	return items, nil
}
