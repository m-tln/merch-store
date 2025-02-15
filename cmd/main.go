package main

import (
	"fmt"
	"log"
	openapi "merch-store/api/generated/go"
	"merch-store/api/handlers"
	"net/http"
)

func main() {
	apiService := handlers.NewCustomAPIService()
	controller := openapi.NewDefaultAPIController(apiService)

	router := openapi.NewRouter(controller)

	fmt.Println("Start listening")
	log.Fatal(http.ListenAndServe(":8080", router))
}
