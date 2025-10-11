package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"freelance-connect-backend/config"
	"freelance-connect-backend/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	routes.TeamRoutes()

	port := os.Getenv("PORT")
	fmt.Println("Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
