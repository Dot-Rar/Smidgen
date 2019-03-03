package main

import (
	"Smidgen/config"
	"Smidgen/database"
	"Smidgen/http"
	"math/rand"
	"time"
)

func main() {
	config.LoadConfig()

	rand.Seed(time.Now().UTC().UnixNano())

	database.Connect()
	database.CreateTables()

	http.LoadTemplates()
	http.StartServer()
}
