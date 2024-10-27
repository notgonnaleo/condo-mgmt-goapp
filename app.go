package main

import (
	"condoapp/src/router"
	"log"
	"net/http"
)

func main() {
	r := router.SetupRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
