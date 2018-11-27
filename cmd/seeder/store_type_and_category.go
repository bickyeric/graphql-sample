package main

import (
	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
)

// StoreTypeCategorySeed ...
func StoreTypeCategorySeed() error {
	for storetype, categories := range storeTypeCategory {
		row, err := database.Connection.Exec(
			`INSERT INTO store_type(name) VALUES (?)`, storetype,
		)
		if err != nil {
			return err
		}
		id, err := row.LastInsertId()
		if err != nil {
			return err
		}

		for index, category := range categories {
			_, err = database.Connection.Exec(
				`INSERT INTO store_category(id, name, type_id) VALUES (?, ?, ?)`, index+1, category, id,
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

var storeTypeCategory = map[string][]string{
	"Beauty & Personal Case": []string{
		"Hair Salon",
		"Message/Therapy",
		"Nail Salon",
		"Spa",
		"Other",
	},
	"Charities & Education": []string{
		"Charitable Organization",
		"Child Care",
		"Religious Organization",
		"School",
		"Tutoring Center",
		"Other",
	},
	"Food & Beverages": []string{
		"Bakery",
		"Bar",
		"Caterer",
		"Coffee Shop",
		"Food Truck",
		"Quick Service Restaurant",
		"Full Service Restaurant",
		"Other",
	},
	"Health Care & Fitness": []string{
		"Chiropractor",
		"Dentist",
		"Eye Care",
		"Gym",
		"Message/Therapy",
		"Pharmacy",
		"Veterinary Service",
		"Other",
	},
	"Home & Repair": []string{
		"Automotive Service",
		"Cleaning",
		"General Contracting",
		"Moving",
		"Painting",
		"Plumbing",
		"Other",
	},
	"Leisure & Entertainment": []string{
		"Cinemas/Theatres",
		"Music",
		"Sporting Events",
		"Wedding Events",
		"Other",
	},
	"Professional Services": []string{
		"Accounting",
		"Consulting",
		"Design",
		"Legal Services",
		"Marketing/Advertising",
		"Photography",
		"Printing Services",
		"Software Development",
		"Other",
	},
	"Retail": []string{
		"Art",
		"Book Store",
		"Clothing & Accessories",
		"Convenience Store",
		"Electronics",
		"Eyewear",
		"Flowers & Gifts",
		"Furniture",
		"Grocery Store",
		"Jewelry",
		"Office Supply",
		"Pet Store",
		"Specialty Store",
		"Sporting Goods",
		"Other",
	},
	"Other": []string{
		"Transportation",
		"Event/Festivals",
		"Other",
	},
}
