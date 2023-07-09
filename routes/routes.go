package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathas/go-rest-api/controllers"
	"github.com/jonathas/go-rest-api/middleware"
)

func HandleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("PUT")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
