package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/JalalA984/apptrak/internal/models"
	"github.com/JalalA984/apptrak/pkg/config"
)

func ping(res http.ResponseWriter, _ *http.Request) {
	res.Write([]byte("OK"))
}

// Home handler with closure-based dependency injection
func Home(app *config.ApplicationConfig) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			notFound(app, res)
			return
		}

		// apps, err := app.Applications.Latest()
		// if err != nil {
		// 	serverError(app, res, err)
		// 	return
		// }

		// for _, application := range apps {
		// 	fmt.Fprintf(res, "%+v\n", application)
		// }

		files := []string{
			"./internal/templates/base.tmpl.html",
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
func Login(app *config.ApplicationConfig) http.HandlerFunc {
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
func Register(app *config.ApplicationConfig) http.HandlerFunc {
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

func ApplicationView(app *config.ApplicationConfig) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(req.URL.Query().Get("id"))
		if err != nil || id < 1 {
			notFound(app, res)
			return
		}

		application, err := app.Applications.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				notFound(app, res)
			} else {
				serverError(app, res, err)
			}
			return
		}

		// Write the snippet data as a plain-text HTTP response body.
		fmt.Fprintf(res, "%+v", application)

	}
}

// Function creates an application and adds it to database
func ApplicationCreate(app *config.ApplicationConfig) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodPost {
			res.Header().Set("Allow", http.MethodPost)
			clientError(app, res, http.StatusMethodNotAllowed)
			return
		}

		// Dummy values for the application record
		name := "TestApp"
		companyName := "SomeCorp"
		position := "Software Engineer"
		status := "Applied"
		applicationDate := time.Now() // Current time as the application date
		var interviewDate *time.Time  // No interview date provided
		notes := "Excited about the opportunity!"

		// Insert the dummy application record
		id, err := app.Applications.Insert(name, companyName, position, status, applicationDate, interviewDate, notes)
		if err != nil {
			serverError(app, res, err)
			return
		}

		// Redirect to a view page for the newly created application
		http.Redirect(res, req, fmt.Sprintf("/application/view?id=%d", id), http.StatusSeeOther)
	}
}
