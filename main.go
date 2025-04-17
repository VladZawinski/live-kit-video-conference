package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/live-kit-video-conference/internal/di"
)

func main() {
	err := godotenv.Load("development.env")
	if err != nil {
		log.Fatal("Error loading env file")
	}
	services := di.InjectServices()
	log.Println(services)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
