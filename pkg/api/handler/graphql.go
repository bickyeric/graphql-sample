package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bickyeric/garut/pkg/graphql/query"
	"github.com/graphql-go/graphql"
	"github.com/julienschmidt/httprouter"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: query.Root,
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
