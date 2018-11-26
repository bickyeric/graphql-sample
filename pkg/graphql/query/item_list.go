package query

import (
	"github.com/bickyeric/garut/pkg/model"
	"github.com/graphql-go/graphql"
)

var itemList = &graphql.Field{
	Type:        graphql.NewList(model.ItemType),
	Description: "Get item list",
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
		return model.Item{}.All(first, offset)
	},
}
