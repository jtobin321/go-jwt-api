package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jtobin321/go-jwt-api/api/controllers"
	"github.com/jtobin321/go-jwt-api/api/seed"
)

var server = controllers.Server{}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	}

	// Load test data (dev only)
	server.Initialize()
	seed.Load(server.DB)

	// Start server on port specified in .env file
	server.Run(os.Getenv("API_PORT"))
}
