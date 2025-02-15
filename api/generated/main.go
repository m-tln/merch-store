package main

import (
	"log"
	"net/http"

	openapi "merch-store/api/generated/go"
)

func main() {
	log.Printf("Server started")

	DefaultAPIService := openapi.NewDefaultAPIService()
	DefaultAPIController := openapi.NewDefaultAPIController(DefaultAPIService)

	router := openapi.NewRouter(DefaultAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
