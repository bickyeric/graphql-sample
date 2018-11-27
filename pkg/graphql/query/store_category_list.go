package query

import (
	"github.com/bickyeric/garut/pkg/graphql/object"
	"github.com/bickyeric/garut/pkg/model"
	"github.com/graphql-go/graphql"
)

var storeCategoryList = &graphql.Field{
	Type:        graphql.NewList(object.StoreCategoryType),
	Description: "Get store category list",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return model.StoreCategory{}.All()
	},
}
