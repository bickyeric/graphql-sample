package object

import (
	"github.com/bickyeric/graphql-sample/pkg/model"
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
		},
	},
)

func init() {
	StoreCategoryType.AddFieldConfig("type", &graphql.Field{
		Type: StoreTypeType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			sc := p.Source.(model.StoreCategory)
			return sc.Type()
		},
	})
	StoreCategoryType.AddFieldConfig("store", &graphql.Field{
		Type: graphql.NewList(StoreType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			sc := p.Source.(model.StoreCategory)
			return sc.Store()
		},
	})
}
