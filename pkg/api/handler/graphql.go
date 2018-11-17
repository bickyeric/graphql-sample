package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bickyeric/warda/pkg/model"
	"github.com/julienschmidt/httprouter"
)

// Handle ...
func Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	peoples, _ := model.People{}.All()
	response, _ := json.Marshal(peoples)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(response)))
}
