package handlers

import (
	"net/http"
	"text/template"

	"github.com/JalalA984/apptrak/pkg/config"
)

func ping(res http.ResponseWriter, _ *http.Request) {
	res.Write([]byte("OK"))
}

// Closure based dependency injection
func Home(app *config.Application) http.HandlerFunc { // Home handler depends on app config
	return func(res http.ResponseWriter, req *http.Request) { // return the closure based function that is our actual handler function
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
			app.ErrorLog.Print(err.Error())
			http.Error(res, "Internal Server Error", 500)
			return
		}

		err = templateSet.ExecuteTemplate(res, "base", nil)
		if err != nil {
			app.ErrorLog.Print(err.Error())
			http.Error(res, "Internal Server Error", 500)
			return
		}
	}
}

func Login(app *config.Application) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		files := []string{
			"./internal/templates/base.tmpl.html",
			"./internal/templates/login.tmpl.html",
		}

		templateSet, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Print((err.Error()))
			http.Error(res, "Internal Server Error", 500)
			return
		}

		err = templateSet.ExecuteTemplate(res, "base", nil)
		if err != nil {
			app.ErrorLog.Print((err.Error()))
			http.Error(res, "Internal Server Error", 500)
			return
		}
	}
}

func Register(app *config.Application) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		files := []string{
			"./internal/templates/base.tmpl.html",
			"./internal/templates/register.tmpl.html",
		}

		templateSet, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Print((err.Error()))
			http.Error(res, "Internal Server Error", 500)
			return
		}

		err = templateSet.ExecuteTemplate(res, "base", nil)
		if err != nil {
			app.ErrorLog.Print((err.Error()))
			http.Error(res, "Internal Server Error", 500)
			return
		}
	}
}
