package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathas/go-rest-api/database"
	"github.com/jonathas/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Order("name asc").Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var p models.Personality
	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var p models.Personality

	json.NewDecoder(r.Body).Decode(&p)
	
	database.DB.Create(&p)
	
	json.NewEncoder(w).Encode(p)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var p models.Personality
	database.DB.First(&p, id)

	json.NewDecoder(r.Body).Decode(&p)
	
	database.DB.Save(&p)
	
	json.NewEncoder(w).Encode(p)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var p models.Personality
	database.DB.Delete(&p, id)
}
