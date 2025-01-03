package handlers

import (
	"net/http"
	"text/template"

	"github.com/JalalA984/apptrak/pkg/config"
)

func ping(res http.ResponseWriter, _ *http.Request) {
	res.Write([]byte("OK"))
}

// Home handler with closure-based dependency injection
func Home(app *config.Application) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(res, req)
			return
		}

		files := []string{
			"./internal/templates/base.tmpl.html",
			"./internal/templates/components/navbar.tmpl.html",
			"./internal/templates/home.tmpl.html",
		}

		templateSet, err := template.ParseFiles(files...)
		if err != nil {
			serverError(app, res, err) // Use serverError from helpers.go
			return
		}

		err = templateSet.ExecuteTemplate(res, "base", nil)
		if err != nil {
			serverError(app, res, err) // Use serverError from helpers.go
		}
	}
}

// Login handler
func Login(app *config.Application) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		files := []string{
			"./internal/templates/base.tmpl.html",
			"./internal/templates/login.tmpl.html",
		}

		templateSet, err := template.ParseFiles(files...)
		if err != nil {
			serverError(app, res, err) // Use serverError from helpers.go
			return
		}

		err = templateSet.ExecuteTemplate(res, "base", nil)
		if err != nil {
			serverError(app, res, err) // Use serverError from helpers.go
		}
	}
}

// Register handler
func Register(app *config.Application) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		files := []string{
			"./internal/templates/base.tmpl.html",
			"./internal/templates/register.tmpl.html",
		}

		templateSet, err := template.ParseFiles(files...)
		if err != nil {
			serverError(app, res, err) // Use serverError from helpers.go
			return
		}

		err = templateSet.ExecuteTemplate(res, "base", nil)
		if err != nil {
			serverError(app, res, err) // Use serverError from helpers.go
		}
	}
}
