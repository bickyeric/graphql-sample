package object

import (
	"github.com/graphql-go/graphql"
)

// StoreCategoryType ...
var StoreCategoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "StoreCategory",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"type_id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
