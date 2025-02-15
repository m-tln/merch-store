package main

import (
	"fmt"
	"log"
	"merch-store/api/controller"
	openapi "merch-store/api/generated/go"
	"merch-store/api/handlers"
	"merch-store/internal/service"
	"net/http"
)

func main() {
	apiService := handlers.NewCustomAPIService()
	controller := controller.NewCustomAPIController(*apiService, service.NewJWTService("pronin"))
	// controller := openapi.NewDefaultAPIController(apiService)

	router := openapi.NewRouter(controller)

	fmt.Println("Start listening")
	log.Fatal(http.ListenAndServe(":8080", router))
}
