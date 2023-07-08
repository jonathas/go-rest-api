package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathas/go-rest-api/controllers"
)

func HandleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
