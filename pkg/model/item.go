package model

import (
	"time"

	"github.com/bickyeric/garut/pkg/wrapper/database"
	"github.com/graphql-go/graphql"
)

// Item ...
type Item struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
func (it Item) All(first, offset int) ([]Item, error) {
	if offset < 1 {
		offset = 20
	}
	var items []Item
	rows, err := database.Connection.Query(
		`SELECT id, name, price, description, image, created_at, updated_at
		FROM items LIMIT ?,?`, first, offset,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i Item
		err := rows.Scan(&i.ID, &i.Name, &i.Price, &i.Description, &i.Image, &i.CreatedAt, &i.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	return items, nil
}
