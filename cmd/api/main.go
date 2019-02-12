package main

import (
	"fmt"
	"log"
	"net/http"

	controllers "github.com/gdlroutes/api/internal/api/controllers/geodata"
	"github.com/gdlroutes/api/internal/api/routers"
	usecases "github.com/gdlroutes/api/internal/api/usecases/geodata"
)

const port = 8080

func main() {
	useCases, err := usecases.New()
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

	log.Printf("Listening on %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
