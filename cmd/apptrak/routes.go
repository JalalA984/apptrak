package main

import (
	"net/http"

	"github.com/JalalA984/apptrak/internal/handlers"
)

func (appConfig *applicationConf) routes() *http.ServeMux {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./public/"))
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	mux.HandleFunc("/", handlers.Home(appConfig.ApplicationConfig))
	mux.HandleFunc("/login", handlers.Login(appConfig.ApplicationConfig))
	mux.HandleFunc("/register", handlers.Register(appConfig.ApplicationConfig))

	mux.HandleFunc("/application/view", handlers.ApplicationView(appConfig.ApplicationConfig))
	mux.HandleFunc("/application/create", handlers.ApplicationCreate(appConfig.ApplicationConfig))

	return mux

}
