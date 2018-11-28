package main

import (
	"log"

	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
	"github.com/bickyeric/graphql-sample/pkg/wrapper/environment"
)

func init() {
	environment.Load()
	database.Connect()
}

func main() {
	log.Println("seeder started...")

	for _, name := range seedNames {
		log.Println(name + "...")
		err := funcs[name]()
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("seeder finished")
}

var seedNames = []string{
	"store_type_and_store_category",
	"store",
	"outlet",
	"employee",
	"item_category",
	"item",
}

var funcs = map[string](func() error){
	"store_type_and_store_category": storeTypeCategorySeed,
	"store":                         storeSeed,
	"outlet":                        outletSeed,
	"employee":                      employeeSeed,
	"item_category":                 itemCategorySeed,
	"item":                          itemSeed,
}
