package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bickyeric/graphql-sample/pkg/wrapper/environment"
	"github.com/nullbio/mig"
)

func main() {
	cmd := os.Args[1]
	log.Println("migration started...")

	environment.Load()

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", environment.DATABASEUSERNAME, environment.DATABASEPASSWORD, environment.DATABASEHOST, environment.DATABASEPORT, environment.DATABASENAME)

	var err error
	if cmd == "up" {
		_, err = mig.Up("mysql", conn, "migration")
	} else if cmd == "down" {
		_, err = mig.DownAll("mysql", conn, "migration")
	}

	if err != nil {
		panic(err)
	}

	log.Println("migration finished")

	os.Exit(0)
}
