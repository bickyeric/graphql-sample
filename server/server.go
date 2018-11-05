package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bickyeric/scaling-chainsaw/route"
)

func Start() {

	s := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		ReadTimeout:    8 * time.Second,
		WriteTimeout:   8 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        route.Web(),
	}

	log.Println("Starting at " + os.Getenv("PORT"))
	log.Fatal(s.ListenAndServe())
}
