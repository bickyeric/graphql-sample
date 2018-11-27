package model

import (
	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
	"github.com/graphql-go/graphql"
)

// Customer ...
type Customer struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Gender      bool   `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
	Comment     string `json:"comment"`
}

// CustomerType ...
var CustomerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Customer",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"first_name": &graphql.Field{
				Type: graphql.String,
			},
			"last_name": &graphql.Field{
				Type: graphql.String,
			},
			"gender": &graphql.Field{
				Type: graphql.Boolean,
			},
			"phone_number": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"city": &graphql.Field{
				Type: graphql.String,
			},
			"zip": &graphql.Field{
				Type: graphql.String,
			},
			"comment": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// All ...
func (c Customer) All(first, offset int) ([]Customer, error) {
	if offset < 1 {
		offset = 20
	}
	var customers []Customer
	rows, err := database.Connection.Query(
		`SELECT id, first_name, last_name, gender, phone_number, address, city, zip, comment
		FROM customers LIMIT ?,?`, first, offset,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Gender, &c.PhoneNumber, &c.Address, &c.City, &c.Zip, &c.Comment)
		customers = append(customers, c)
	}

	return customers, nil
}

// GetByID ...
func (c Customer) GetByID() (*Customer, error) {
	row := database.Connection.QueryRow(
		`SELECT id, first_name, last_name, gender, phone_number, address, city, zip, comment
		FROM customers WHERE id=?`, c.ID,
	)
	err := row.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Gender, &c.PhoneNumber, &c.Address, &c.City, &c.Zip, &c.Comment)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
