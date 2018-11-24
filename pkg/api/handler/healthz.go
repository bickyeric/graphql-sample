package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Healthz ...
func Healthz(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
