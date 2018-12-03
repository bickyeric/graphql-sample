package query

import (
	"github.com/bickyeric/graphql-sample/pkg/graphql/object"
	"github.com/bickyeric/graphql-sample/pkg/model"
	"github.com/graphql-go/graphql"
)

var employeeList = &graphql.Field{
	Type:        graphql.NewList(object.EmployeeType),
	Description: "Get employee list",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return model.Employee{}.All()
	},
}
