package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/JalalA984/apptrak/internal/handlers"
)

func main() {
	port := flag.String("port", ":5000", "HTTP Network Address")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./public/"))
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/register", handlers.Register)

	log.Printf("Application started on %s", *port)
	err := http.ListenAndServe(*port, mux)
	log.Fatal(err)
}
