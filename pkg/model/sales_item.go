package model

import "github.com/graphql-go/graphql"

// SalesItem ..
type SalesItem struct {
	ID       int    `json:"id"`
	Name     string `json:"status"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantiy"`
}

// SalesItemType ...
var SalesItemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SalesItem",
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
			"quantity": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
