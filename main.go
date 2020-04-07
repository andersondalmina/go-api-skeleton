package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andersondalmina/go-api-skeleton/config"
	"github.com/gorilla/context"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	r := config.Handlers()

	log.Printf("Listening at :%s", os.Getenv("API_PORT"))
	err := http.ListenAndServe(":"+os.Getenv("API_PORT"), context.ClearHandler(r))
	if err != nil {
		log.Fatal(err)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
}
