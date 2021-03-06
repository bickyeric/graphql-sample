package query

import (
	"github.com/bickyeric/graphql-sample/pkg/graphql/object"
	"github.com/bickyeric/graphql-sample/pkg/model"
	"github.com/graphql-go/graphql"
)

var outletList = &graphql.Field{
	Type:        graphql.NewList(object.OutletType),
	Description: "Get outlet list",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return model.Outlet{}.All()
	},
}
