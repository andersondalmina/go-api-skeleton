package config

import (
	"os"
	"time"

	"github.com/andersondalmina/go-api-skeleton/handlers"
	"github.com/andersondalmina/go-api-skeleton/infrastructure"
	"github.com/andersondalmina/go-api-skeleton/models"
	"github.com/andersondalmina/go-api-skeleton/security"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Handlers configure the routes and handlers
func Handlers() *mux.Router {
	jwtm := security.CreateJWTManager(getJWTOptions())

	db := infrastructure.CreateDatabase()

	// Configure repositories
	uR := models.NewUserRepository(db)

	router := mux.NewRouter().StrictSlash(true)
	v1Router := router.PathPrefix("/v1").Subrouter()

	v1Router.HandleFunc("/", handlers.HomeHandler()).Methods("GET")
	v1Router.HandleFunc("/login", handlers.LoginHandler(uR, jwtm)).Methods("POST")
	v1Router.HandleFunc("/register", handlers.RegisterHandler(uR, jwtm)).Methods("POST")

	// Authenticated routes
	verifyAuthenticationHandler := negroni.New()
	verifyAuthenticationHandler.UseFunc(handlers.VerifyAuthenticationHandler(jwtm))

	v1Router.Path("/profile").Handler(verifyAuthenticationHandler.With(
		negroni.WrapFunc(handlers.ProfileHandler(uR)),
	)).Methods("GET")

	return router
}

func getJWTOptions() security.JWTOptions {
	return security.JWTOptions{
		SigningMethod:  os.Getenv("SIGNING_METHOD"),
		PrivateKeyPath: os.Getenv("PRIVATE_KEY"), // $ openssl genrsa -out app.rsa keysize
		PublicKeyPath:  os.Getenv("PUBLIC_KEY"),  // $ openssl rsa -in app.rsa -pubout > app.rsa.pub
		Expiration:     60 * time.Minute,
	}
}
