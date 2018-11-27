package query

import (
	"github.com/graphql-go/graphql"
)

// Root ...
var Root = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"customerList":      customerList,
			"itemList":          itemList,
			"outletList":        outletList,
			"salesList":         salesList,
			"storeList":         storeList,
			"storeCategoryList": storeCategoryList,
			"storeTypeList":     storeTypeList,
		},
	},
)
