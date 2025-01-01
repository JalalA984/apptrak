package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/JalalA984/apptrak/internal/handlers"
)

func main() {
	port := flag.String("port", ":5000", "HTTP Network Address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./public/"))
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/register", handlers.Register)

	infoLog.Printf("Application started on %s", *port)
	err := http.ListenAndServe(*port, mux)
	errorLog.Fatal(err)
}
