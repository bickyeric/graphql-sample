package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bickyeric/garut/pkg/wrapper/environment"
	"github.com/nullbio/mig"
)

func main() {
	log.Println("migration started")

	environment.Load()

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", environment.DATABASEUSERNAME, environment.DATABASEPASSWORD, environment.DATABASEHOST, environment.DATABASEPORT, environment.DATABASENAME)

	_, err := mig.Up("mysql", conn, "migration")
	if err != nil {
		panic(err)
	}

	log.Println("migration finished")

	os.Exit(0)
}
