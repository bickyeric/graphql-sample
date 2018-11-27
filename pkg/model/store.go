package model

import (
	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
)

// Store ...
type Store struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	TypeID      int    `json:"type_id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Twitter     string `json:"twitter"`
	Facebook    string `json:"facebook"`
	Instagram   string `json:"instagram"`
	Image       string `json:"image"`
}

// All ...
func (s Store) All(first, offset int) ([]Store, error) {
	if offset < 1 {
		offset = 20
	}
	var stores []Store
	rows, err := database.Connection.Query(
		`SELECT id, category_id, type_id, name, address, city, state, zip, email, description, website, twitter, facebook, instagram, image
		FROM store LIMIT ?,?`, first, offset,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s Store
		err := rows.Scan(&s.ID, &s.CategoryID, &s.TypeID, &s.Name, &s.Address, &s.City, &s.State, &s.Zip, &s.Email, &s.Description, &s.Website, &s.Twitter, &s.Facebook, &s.Instagram, &s.Image)
		if err != nil {
			return nil, err
		}
		stores = append(stores, s)
	}

	return stores, nil
}

// Category ...
func (s *Store) Category() (StoreCategory, error) {
	return StoreCategory{ID: s.CategoryID}.GetByID()
}

// Type ...
func (s *Store) Type() (StoreType, error) {
	return StoreType{ID: s.TypeID}.GetByID()
}

// GetByID ...
func (s Store) GetByID() (Store, error) {
	row := database.Connection.QueryRow(
		`SELECT id, category_id, type_id, name, address, city, state, zip, email, description, website, twitter, facebook, instagram, image
		FROM store WHERE id=?`, s.ID,
	)
	err := row.Scan(&s.ID, &s.CategoryID, &s.TypeID, &s.Name, &s.Address, &s.City, &s.State, &s.Zip, &s.Email, &s.Description, &s.Website, &s.Twitter, &s.Facebook, &s.Instagram, &s.Image)
	if err != nil {
		return s, err
	}
	return s, nil
}

// Outlet ...
func (s *Store) Outlet() ([]Outlet, error) {
	return Outlet{StoreID: s.ID}.GetByStoreID()
}
