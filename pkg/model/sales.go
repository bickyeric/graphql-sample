package model

import (
	"time"

	"github.com/bickyeric/garut/pkg/wrapper/database"
	"github.com/graphql-go/graphql"
)

// Sales ...
type Sales struct {
	ID         int       `json:"id"`
	Status     string    `json:"status"`
	TotalPrice int       `json:"total_price"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CustomerID int       `json:"customer_id"`
}

// SalesType ...
var SalesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Sales",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"total_price": &graphql.Field{
				Type: graphql.Int,
			},
			"comment": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"customer": &graphql.Field{
				Type: CustomerType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					sales := p.Source.(Sales)
					return sales.Customer()
				},
			},
			"items": &graphql.Field{
				Type: graphql.NewList(SalesItemType),
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					first, _ := p.Args["first"].(int)
					offset, _ := p.Args["offset"].(int)
					sales := p.Source.(Sales)
					return sales.Items(first, offset)
				},
			},
		},
	},
)

// All ...
func (s Sales) All(first, offset int) ([]Sales, error) {
	if offset < 1 {
		offset = 20
	}
	var sales []Sales
	rows, err := database.Connection.Query(
		`SELECT id, total_price, status, comment, created_at, updated_at, customer_id
		FROM sales LIMIT ?,?`, first, offset,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&s.ID, &s.TotalPrice, &s.Status, &s.Comment, &s.CreatedAt, &s.UpdatedAt, &s.CustomerID)
		sales = append(sales, s)
	}

	return sales, nil
}

// Customer ...
func (s *Sales) Customer() (*Customer, error) {
	return Customer{ID: s.CustomerID}.GetByID()
}

// Items ...
func (s *Sales) Items(first, offset int) ([]SalesItem, error) {
	if offset < 1 {
		offset = 20
	}
	var items []SalesItem
	rows, err := database.Connection.Query(
		`SELECT items.id, items.name, sales_detail.price, sales_detail.quantity_purchased
		FROM items, sales_detail
		WHERE items.id=sales_detail.item_id LIMIT ?,?`, first, offset,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i SalesItem
		err := rows.Scan(&i.ID, &i.Name, &i.Price, &i.Quantity)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}
