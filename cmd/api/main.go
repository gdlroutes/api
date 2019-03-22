package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	geodataController "github.com/gdlroutes/api/internal/api/controllers/geodata"
	userController "github.com/gdlroutes/api/internal/api/controllers/user"
	"github.com/gdlroutes/api/internal/api/middleware"
	"github.com/gdlroutes/api/internal/api/routers"
	geodataUsecases "github.com/gdlroutes/api/internal/api/usecases/geodata"
	geodataStorage "github.com/gdlroutes/api/internal/api/usecases/geodata/storage"
	userUsecases "github.com/gdlroutes/api/internal/api/usecases/user"
	userStorage "github.com/gdlroutes/api/internal/api/usecases/user/storage"
	userToken "github.com/gdlroutes/api/internal/api/usecases/user/token"
	"github.com/justinas/alice"
)

const (
	portEnvName = "PORT"
	defaultPort = "8080"

	cookieDomainEnvName = "COOKIE_DOMAIN"
	defaultCookieDomain = "localhost"

	corsOriginEnvName = "CORS_ORIGIN"
	defaultCORSOrigin = "*"

	tokenDurationEnvName = "TOKEN_DURATION"
	defaultTokenDuration = "24h"

	tokenKeyEnvName = "TOKEN_KEY"
	defaultTokenKey = "1ns3cur3"
)

var (
	port          string
	corsOrigin    string
	cookieDomain  string
	tokenDuration string
	tokenKey      string
)

func init() {
	port = os.Getenv(portEnvName)
	if port == "" {
		port = defaultPort
	}

	corsOrigin = os.Getenv(corsOriginEnvName)
	if corsOrigin == "" {
		corsOrigin = defaultCORSOrigin
	}

	cookieDomain = os.Getenv(cookieDomainEnvName)
	if cookieDomain == "" {
		cookieDomain = defaultCookieDomain
	}

	tokenDuration = os.Getenv(tokenDurationEnvName)
	if tokenDuration == "" {
		tokenDuration = defaultTokenDuration
	}

	tokenKey = os.Getenv(tokenKeyEnvName)
	if tokenKey == "" {
		log.Println("WARNING: using default token-signing key")
		tokenKey = defaultTokenKey
	}
}

func main() {

	// Geodata
	geodataStorage, err := geodataStorage.NewFake()
	if err != nil {
		log.Fatalf("error creating geodata storage: %v", err)
	}
	geodataUseCases, err := geodataUsecases.New(geodataStorage)
	if err != nil {
		log.Fatalf("error creating geodata usecases: %v", err)
	}
	geodataController, err := geodataController.New(geodataUseCases)
	if err != nil {
		log.Fatal("error creating geodata controller", err)
	}
	geodataRouter := &routers.Hotspot{Controller: geodataController}

	// User
	userStorage, err := userStorage.NewFake()
	if err != nil {
		log.Fatalf("error creating user storage: %v", err)
	}
	duration, err := time.ParseDuration(tokenDuration)
	if err != nil {
		log.Fatalf("error parsing token duration: %v", err)
	}
	userTokenGenerator, err := userToken.NewJWT(duration, tokenKey)
	if err != nil {
		log.Fatalf("error creating user token generator: %v", err)
	}
	userUseCases, err := userUsecases.New(userStorage, userTokenGenerator)
	if err != nil {
		log.Fatalf("error creating user usecases: %v", err)
	}
	userController, err := userController.New(userUseCases, cookieDomain)
	if err != nil {
		log.Fatalf("error creating user controller: %v", err)
	}
	userRouter := &routers.User{Controller: userController}

	// Main router
	mux := http.NewServeMux()
	mux.Handle(routers.GeodataPrefix, geodataRouter)
	mux.Handle(routers.UserPrefix, userRouter)

	// Chaining middlewares
	server := alice.New(middleware.CORS(corsOrigin), middleware.Token()).Then(mux)

	log.Printf("Listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}
