package model

import (
	"database/sql"
	"time"

	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
	"github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
)

// Customer ...
type Customer struct {
	ID          int `json:"id"`
	StoreID     int
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"first_name"`
	Email       string
	Birthday    time.Time
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string
	Zip         string `json:"zip"`

	phoneNumber sql.NullString
	email       sql.NullString
	birthday    mysql.NullTime
	address     sql.NullString
	city        sql.NullString
	state       sql.NullString
	zip         sql.NullString
}

// Create ...
func (c Customer) Create() (Customer, error) {
	row, err := database.Connection.Exec(
		`INSERT INTO customer(store_id, phone_number, name, email, birthday, address, city, state, zip)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`, c.StoreID, NewNullString(c.PhoneNumber), c.Name, NewNullString(c.Email), NewNullTime(c.Birthday), NewNullString(c.Address), NewNullString(c.City), NewNullString(c.State), NewNullString(c.Zip),
	)
	if err != nil {
		return c, err
	}
	id, _ := row.LastInsertId()
	c.ID = int(id)
	return c, nil
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
	// if offset < 1 {
	// 	offset = 20
	// }
	var customers []Customer
	// rows, err := database.Connection.Query(
	// 	`SELECT id, first_name, last_name, gender, phone_number, address, city, zip, comment
	// 	FROM customers LIMIT ?,?`, first, offset,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// for rows.Next() {
	// 	rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Gender, &c.PhoneNumber, &c.Address, &c.City, &c.Zip, &c.Comment)
	// 	customers = append(customers, c)
	// }

	return customers, nil
}

// GetByID ...
func (c Customer) GetByID() (*Customer, error) {
	// row := database.Connection.QueryRow(
	// 	`SELECT id, first_name, last_name, gender, phone_number, address, city, zip, comment
	// 	FROM customers WHERE id=?`, c.ID,
	// )
	// err := row.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Gender, &c.PhoneNumber, &c.Address, &c.City, &c.Zip, &c.Comment)
	// if err != nil {
	// 	return nil, err
	// }
	return &c, nil
}

// GetByStoreID ...
func (c Customer) GetByStoreID() ([]Customer, error) {
	var customers []Customer
	rows, err := database.Connection.Query(
		`SELECT id, store_id, phone_number, name, email, birthday, address, city, state, zip
		FROM customer
		WHERE store_id=?`, c.StoreID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&c.ID, &c.StoreID, &c.phoneNumber, &c.Name, &c.email, &c.birthday, &c.address, &c.city, &c.state, &c.zip)
		c.fill()
		customers = append(customers, c)
	}

	return customers, nil
}

func (c *Customer) fill() {
	if c.phoneNumber.Valid {
		c.PhoneNumber = c.phoneNumber.String
	}
	if c.email.Valid {
		c.Email = c.email.String
	}
	if c.birthday.Valid {
		c.Birthday = c.birthday.Time
	}
	if c.address.Valid {
		c.Address = c.address.String
	}
	if c.city.Valid {
		c.City = c.city.String
	}
	if c.state.Valid {
		c.State = c.state.String
	}
	if c.zip.Valid {
		c.Zip = c.zip.String
	}
}
