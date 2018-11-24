package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bickyeric/garut/pkg/model"

	"github.com/graphql-go/graphql"
	"github.com/julienschmidt/httprouter"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"salesList": &graphql.Field{
				Type:        graphql.NewList(model.SalesType),
				Description: "Get sales list",
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					first, _ := p.Args["first"].(int)
					offset, _ := p.Args["offset"].(int)
					return model.Sales{}.All(first, offset)
				},
			},

			"customerList": &graphql.Field{
				Type:        graphql.NewList(model.CustomerType),
				Description: "Get customer list",
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					first, _ := p.Args["first"].(int)
					offset, _ := p.Args["offset"].(int)
					return model.Customer{}.All(first, offset)
				},
			},

			"itemList": &graphql.Field{
				Type:        graphql.NewList(model.ItemType),
				Description: "Get item list",
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					first, _ := p.Args["first"].(int)
					offset, _ := p.Args["offset"].(int)
					return model.Item{}.All(first, offset)
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

// GraphQL ...
func GraphQL(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result := executeQuery(r.URL.Query().Get("query"), schema)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
