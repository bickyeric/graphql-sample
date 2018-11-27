package query

import (
	"github.com/bickyeric/garut/pkg/graphql/object"
	"github.com/bickyeric/garut/pkg/model"
	"github.com/graphql-go/graphql"
)

var storeTypeList = &graphql.Field{
	Type:        graphql.NewList(object.StoreTypeType),
	Description: "Get store type list",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return model.StoreType{}.All()
	},
}
