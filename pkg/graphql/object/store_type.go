package object

import (
	"github.com/bickyeric/garut/pkg/model"
	"github.com/graphql-go/graphql"
)

// StoreTypeType ...
var StoreTypeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "StoreType",
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
	StoreTypeType.AddFieldConfig("category", &graphql.Field{
		Type: graphql.NewList(StoreCategoryType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			sc := p.Source.(model.StoreType)
			return model.StoreCategory{TypeID: sc.ID}.GetByTypeID()
		},
	})
}
