package routes

import (
	"go-api-project/controllers/root"
	"go-api-project/controllers/user"

	"github.com/gorilla/mux"
)

// UseRoutes assign contrroller to the routers
func UseRoutes(app *mux.Router) {
	app.HandleFunc("/", root.HomeHandler).Methods("GET")
	app.HandleFunc("/user/{userId}", user.IDHandler).Methods("GET")
}
