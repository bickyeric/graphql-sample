package main

import (
	"log"

	"github.com/bickyeric/garut/pkg/wrapper/database"
	"github.com/bickyeric/garut/pkg/wrapper/environment"
)

func init() {
	environment.Load()
	database.Connect()
}

func main() {
	log.Println("seeder started...")

	for s, fn := range funcs {
		log.Println(s + "...")
		err := fn()
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("seeder finished")
}

var funcs = map[string](func() error){
	"store_type": StoreTypeCategorySeed,
}
