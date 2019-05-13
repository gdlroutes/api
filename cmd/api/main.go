package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/justinas/alice"
	_ "github.com/lib/pq" // Postgres driver

	"github.com/gdlroutes/api/internal/api/middleware"
	"github.com/gdlroutes/api/internal/api/router"
	"github.com/gdlroutes/api/internal/api/usecases/geodata"
	geodatastg "github.com/gdlroutes/api/internal/api/usecases/geodata/storage"
	"github.com/gdlroutes/api/internal/api/usecases/user"
	userstg "github.com/gdlroutes/api/internal/api/usecases/user/storage"
	userToken "github.com/gdlroutes/api/internal/api/usecases/user/token"
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

	fakeDatabaseEnvName = "FAKE_DB"

	databaseHostEnvName = "DB_HOST"
	defaultDatabaseHost = "127.0.0.1"

	databasePortEnvName = "DB_PORT"
	defaultDatabasePort = "5432"

	databaseDBEnvName = "DB_DATABASE"
	defaultDatabaseDB = "gdlroutes"

	databaseUserEnvName = "DB_USER"
	defaultDatabaseUser = "postgres"

	databasePasswordEnvName = "DB_PASSWORd"
	defaultDatabasePassword = "postgres"

	databaseURLEnvName = "DATABASE_URL"
)

var (
	port          string
	corsOrigin    string
	cookieDomain  string
	tokenDuration string
	tokenKey      string
	databaseURL   string
	fakeDatabase  bool
	db            *sql.DB
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

	if os.Getenv(fakeDatabaseEnvName) != "" {
		fakeDatabase = true
		return
	}

	databaseURL = os.Getenv(databaseURLEnvName)
	if databaseURL == "" {
		host := os.Getenv(databaseHostEnvName)
		if host == "" {
			host = defaultDatabaseHost
		}
		port := os.Getenv(databasePortEnvName)
		if port == "" {
			port = defaultDatabasePort
		}
		db := os.Getenv(databaseDBEnvName)
		if db == "" {
			db = defaultDatabaseDB
		}
		user := os.Getenv(databaseUserEnvName)
		if user == "" {
			user = defaultDatabaseUser
		}
		password := os.Getenv(databasePasswordEnvName)
		if password == "" {
			password = defaultDatabasePassword
		}
		ssl := url.Values{}
		ssl.Set("sslmode", "disable")

		dsn := url.URL{
			Scheme:   "postgres",
			User:     url.UserPassword(user, password),
			Host:     fmt.Sprintf("%s:%s", host, port),
			Path:     db,
			RawQuery: ssl.Encode(),
		}
		databaseURL = dsn.String()
	}

	var err error
	db, err = sql.Open("postgres", databaseURL)
	if err != nil {
		panic(fmt.Sprintf("unable to create postgres database driver: %v", err))

	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("unable to connect: %v", err))
	}
}

func main() {
	var (
		geodataStorage geodata.Storage
		userStorage    user.Storage
		err            error
	)

	// Geodata
	if fakeDatabase {
		geodataStorage, err = geodatastg.NewFake()
		if err != nil {
			log.Fatalf("error creating fake geodata storage: %v", err)
		}
	} else {
		geodataStorage, err = geodatastg.NewPostgres(db)
		if err != nil {
			log.Fatalf("error creating Postgres geodata storage: %v", err)
		}
	}
	geodataUseCases, err := geodata.New(geodataStorage)
	if err != nil {
		log.Fatalf("error creating geodata usecases: %v", err)
	}

	// User
	if fakeDatabase {
		userStorage, err = userstg.NewFake()
		if err != nil {
			log.Fatalf("error creating fake user storage: %v", err)
		}
	} else {
		userStorage, err = userstg.NewPostgres(db)
		if err != nil {
			log.Fatalf("error creating Postgres user storage: %v", err)
		}
	}
	duration, err := time.ParseDuration(tokenDuration)
	if err != nil {
		log.Fatalf("error parsing token duration: %v", err)
	}
	userTokenGenerator, err := userToken.NewJWT(duration, tokenKey)
	if err != nil {
		log.Fatalf("error creating user token generator: %v", err)
	}
	userUseCases, err := user.New(userStorage, userTokenGenerator)
	if err != nil {
		log.Fatalf("error creating user usecases: %v", err)
	}

	// Main router
	router := &router.Router{
		GeodataUseCases: geodataUseCases,
		UserUseCases:    userUseCases,
		CookieDomain:    cookieDomain,
	}

	// Chaining middlewares
	server := alice.New(middleware.RequestLogger(), middleware.CORS(corsOrigin), middleware.Token()).Then(router)

	log.Printf("Listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}
