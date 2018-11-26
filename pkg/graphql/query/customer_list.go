package query

import (
	"github.com/bickyeric/garut/pkg/model"
	"github.com/graphql-go/graphql"
)

var customerList = &graphql.Field{
	Type:        graphql.NewList(model.CustomerType),
	Description: "Get customer list",
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
		return model.Customer{}.All(first, offset)
	},
}
