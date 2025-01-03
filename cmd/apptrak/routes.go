package main

import (
	"net/http"

	"github.com/JalalA984/apptrak/internal/handlers"
)

func (appConfig *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./public/"))
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	mux.HandleFunc("/", handlers.Home(appConfig.Application))
	mux.HandleFunc("/login", handlers.Login(appConfig.Application))
	mux.HandleFunc("/register", handlers.Register(appConfig.Application))

	return mux

}
