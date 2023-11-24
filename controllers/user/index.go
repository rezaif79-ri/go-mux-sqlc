package user

import (
	"encoding/json"
	"go-api-project/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// IDHandler handle response with method GET
func IDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response := utils.CreateResponseOk("", nil)

	data := make(map[string]interface{})
	data["userId"] = vars["userId"]
	response.Data = data

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
