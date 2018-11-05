package route

import (
	"github.com/bickyeric/scaling-chainsaw/handler/api/healthz"
	"github.com/julienschmidt/httprouter"
)

func Web() *httprouter.Router {
	var router *httprouter.Router = httprouter.New()

	router.HandlerFunc("GET", "/healthz", healthz.Handle)

	return router
}
