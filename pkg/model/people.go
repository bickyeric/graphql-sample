package model

import (
	"github.com/bickyeric/warda/pkg/wrapper/database"
)

// People ...
type People struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Gender      bool   `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address1    string `json:"address_1"`
	Address2    string `json:"address_2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Country     string `json:"country"`
	Comment     string `json:"comments"`
}

// All ...
func (p People) All() ([]People, error) {
	var peoples []People
	rows, err := database.Connection.Query(
		`SELECT id, first_name, last_name, gender, phone_number, email, address_1, address_2, city, state, zip, country, comments
		FROM people`,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var people People
		rows.Scan(&people.ID, &people.FirstName, &people.LastName, &people.Gender, &people.PhoneNumber, &people.Email, &people.Address1, &people.Address2, &people.City, &people.State, &people.Zip, &people.Country, &people.Comment)
		peoples = append(peoples, people)
	}

	return peoples, nil
}
