package object

import (
	"github.com/bickyeric/graphql-sample/pkg/model"
	"github.com/graphql-go/graphql"
)

// StoreType ...
var StoreType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Store",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"city": &graphql.Field{
				Type: graphql.String,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
			"zip": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"website": &graphql.Field{
				Type: graphql.String,
			},
			"twitter": &graphql.Field{
				Type: graphql.String,
			},
			"facebook": &graphql.Field{
				Type: graphql.String,
			},
			"instagram": &graphql.Field{
				Type: graphql.String,
			},
			"image": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func init() {
	StoreType.AddFieldConfig("outlet", &graphql.Field{
		Type: graphql.NewList(OutletType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			s := p.Source.(model.Store)
			return s.Outlet()
		},
	})
	StoreType.AddFieldConfig("type", &graphql.Field{
		Type: StoreTypeType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			s := p.Source.(model.Store)
			return s.Type()
		},
	})
	StoreType.AddFieldConfig("category", &graphql.Field{
		Type: StoreCategoryType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			s := p.Source.(model.Store)
			return s.Category()
		},
	})
}
