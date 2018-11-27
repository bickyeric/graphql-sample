package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bickyeric/graphql-sample/pkg/api/handler"
	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
	"github.com/bickyeric/graphql-sample/pkg/wrapper/environment"

	"github.com/julienschmidt/httprouter"
)

// Run ...
func Run() {

	environment.Load()

	database.Connect()

	router := httprouter.New()
	router.GET("/graphql", handler.GraphQL)
	router.GET("/healthz", handler.Healthz)

	if !environment.IsProduction() {
		log.Println(fmt.Sprintf("port: %s", environment.APIPORT))
	}

	log.Fatal(http.ListenAndServe(":1234", router))
}
