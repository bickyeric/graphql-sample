package query

import (
	"github.com/bickyeric/graphql-sample/pkg/graphql/object"
	"github.com/bickyeric/graphql-sample/pkg/model"
	"github.com/graphql-go/graphql"
)

var outletList = &graphql.Field{
	Type:        graphql.NewList(object.OutletType),
	Description: "Get outlet list",
	Args: graphql.FieldConfigArgument{
		"first": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"offset": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// first, _ := p.Args["first"].(int)
		// offset, _ := p.Args["offset"].(int)
		return model.Outlet{}.All()
	},
}
