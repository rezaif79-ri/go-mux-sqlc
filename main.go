package main

import (
	_ "go-api-project/config"
	"go-api-project/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Logger is a middleware to show traffic request
func Logger(handler http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, handler)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := mux.NewRouter().StrictSlash(true)
	app.Use(Logger)
	routes.UseRoutes(app)

	log.Fatal(http.ListenAndServe("127.0.0.1:"+os.Getenv("PORT"), app))
}
