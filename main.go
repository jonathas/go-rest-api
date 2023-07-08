package main

import (
	"github.com/jonathas/go-rest-api/models"
	"github.com/jonathas/go-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Jonathas", History: "Jonathas is a software developer."},
		{Id: 2, Name: "Jon", History: "Jon is a software developer."},
	}

	routes.HandleRequests()
}