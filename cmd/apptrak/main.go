package main

import (
	"log"
	"net/http"
	"github.com/JalalA984/apptrak/internal/handlers"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)

	log.Print("Application started on ", port)

	err := http.ListenAndServe(port, mux)

	log.Fatal(err)
}
