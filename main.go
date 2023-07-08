package main

import (
	"github.com/jonathas/go-rest-api/database"
	"github.com/jonathas/go-rest-api/routes"
)

func main() {
	database.Init()
	routes.HandleRequests()
}