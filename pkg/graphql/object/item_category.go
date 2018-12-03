package object

import (
	"github.com/graphql-go/graphql"
)

// ItemCategoryType ...
var ItemCategoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ItemCategory",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func init() {
}
