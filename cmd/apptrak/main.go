package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/JalalA984/apptrak/internal/handlers"
	"github.com/JalalA984/apptrak/pkg/config"
)

func main() {
	port := flag.String("port", ":5000", "HTTP Network Address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Llongfile)

	appConfig := &config.Application{
		ErrorLog: errorLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./public/"))
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	mux.HandleFunc("/", handlers.Home(appConfig))
	mux.HandleFunc("/login", handlers.Login(appConfig))
	mux.HandleFunc("/register", handlers.Register(appConfig))

	server := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Application starting on %s", *port)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
