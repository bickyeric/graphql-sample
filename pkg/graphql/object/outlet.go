package object

import (
	"github.com/bickyeric/graphql-sample/pkg/model"
	"github.com/graphql-go/graphql"
)

// OutletType ...
var OutletType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Outlet",
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
			"phone_number": &graphql.Field{
				Type: graphql.String,
			},
			"city": &graphql.Field{
				Type: graphql.String,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.Boolean,
			},
			"zip": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func init() {
	OutletType.AddFieldConfig("store", &graphql.Field{
		Type: StoreType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			o := p.Source.(model.Outlet)
			return o.Store()
		},
	})
}
