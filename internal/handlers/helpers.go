package handlers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/JalalA984/apptrak/pkg/config"
)

// serverError logs the error and sends a 500 Internal Server Error response.
func serverError(app *config.ApplicationConfig, res http.ResponseWriter, err error) {
	stackTrace := fmt.Sprintf("[SERVER ERROR]: %s\n[STACK TRACE]:\n%s\n", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, stackTrace)

	http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func clientError(app *config.ApplicationConfig, res http.ResponseWriter, status int) {
	http.Error(res, http.StatusText(status), status)
}

func notFound(app *config.ApplicationConfig, res http.ResponseWriter) {
	clientError(app, res, http.StatusNotFound)
}
