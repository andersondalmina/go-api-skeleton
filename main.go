package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andersondalmina/go-api-skeleton/handlers"
	"github.com/andersondalmina/go-api-skeleton/infrastructure"
	"github.com/andersondalmina/go-api-skeleton/models"
	"github.com/andersondalmina/go-api-skeleton/security"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/urfave/negroni"
)

func main() {
	loadEnv()

	jwtm := security.CreateJWTManager(getJWTOptions())

	db := infrastructure.CreateDatabase()

	uR := models.NewUserRepository(db)

	authenticateMiddleware := negroni.New()
	authenticateMiddleware.UseFunc(handlers.AuthenticateMiddleware(jwtm))

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler(uR, jwtm)).Methods("POST")
	r.HandleFunc("/register", handlers.RegisterHandler(uR)).Methods("POST")

	r.Path("/teste").Handler(authenticateMiddleware.With(
		negroni.WrapFunc(handlers.HomeHandler),
	)).Methods("GET")

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

func getJWTOptions() security.JWTOptions {
	return security.JWTOptions{
		SigningMethod:  os.Getenv("SIGNING_METHOD"),
		PrivateKeyPath: os.Getenv("PRIVATE_KEY"), // $ openssl genrsa -out app.rsa keysize
		PublicKeyPath:  os.Getenv("PUBLIC_KEY"),  // $ openssl rsa -in app.rsa -pubout > app.rsa.pub
		Expiration:     60 * time.Minute,
	}
}
