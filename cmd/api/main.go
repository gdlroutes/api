package main

import (
	"fmt"
	"log"
	"net/http"

	controllers "github.com/gdlroutes/api/internal/api/controllers/geodata"
	"github.com/gdlroutes/api/internal/api/routers"
	usecases "github.com/gdlroutes/api/internal/api/usecases/geodata"
	"github.com/gdlroutes/api/internal/api/usecases/geodata/storage"
)

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		h.ServeHTTP(w, r)
	})
}

const port = 8080

func main() {
	storage, err := storage.NewFake()
	if err != nil {
		log.Fatalf("error creating geodata storage: %v", err)
	}
	useCases, err := usecases.New(storage)
	if err != nil {
		log.Fatalf("error creating geodata usecases: %v", err)
	}
	controller, err := controllers.New(useCases)
	if err != nil {
		log.Fatal("error creating geodata controller", err)
	}
	router := &routers.Hotspot{Controller: controller}

	mux := http.NewServeMux()
	mux.Handle(routers.GeodataPrefix, router)

	server := cors(mux)

	log.Printf("Listening on %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}
