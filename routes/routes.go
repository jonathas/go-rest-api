package routes

import (
	"log"
	"net/http"

	"github.com/jonathas/go-rest-api/controllers"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
