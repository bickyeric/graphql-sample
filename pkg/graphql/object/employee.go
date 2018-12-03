package object

import (
	"github.com/bickyeric/graphql-sample/pkg/model"
	"github.com/graphql-go/graphql"
)

// EmployeeType ...
var EmployeeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Employee",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"first_name": &graphql.Field{
				Type: graphql.String,
			},
			"last_name": &graphql.Field{
				Type: graphql.String,
			},
			"phone_number": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"confirmed": &graphql.Field{
				Type: graphql.Boolean,
			},
			"active": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

func init() {
	EmployeeType.AddFieldConfig("outlet", &graphql.Field{
		Type: OutletType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			e := p.Source.(model.Employee)
			return e.Outlet()
		},
	})
	EmployeeType.AddFieldConfig("store", &graphql.Field{
		Type: StoreType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			e := p.Source.(model.Employee)
			return e.Store()
		},
	})
}
