package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/JalalA984/apptrak/pkg/config"
)

// Define a wrapper around config.Application
type application struct {
	*config.Application
}

func main() {
	port := flag.String("port", ":5000", "HTTP Network Address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Llongfile)

	appConfig := &application{
		Application: &config.Application{
			ErrorLog: errorLog,
		},
	}

	server := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  appConfig.routes(),
	}

	infoLog.Printf("Application starting on %s", *port)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
