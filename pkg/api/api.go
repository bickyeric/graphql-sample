package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bickyeric/warda/pkg/api/handler"
	"github.com/bickyeric/warda/pkg/wrapper/database"
	"github.com/bickyeric/warda/pkg/wrapper/environment"
	"github.com/julienschmidt/httprouter"
)

// Run ...
func Run() {

	environment.Load()

	database.Connect()

	router := httprouter.New()
	router.GET("/graphql", handler.Handle)
	router.GET("/healthz", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	if !environment.IsProduction() {
		log.Println(fmt.Sprintf("port: %s", environment.APIPORT))
	}

	log.Fatal(http.ListenAndServe(":1234", router))
}
