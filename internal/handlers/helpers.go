package handlers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/JalalA984/apptrak/pkg/config"
)

// serverError logs the error and sends a 500 Internal Server Error response.
func serverError(app *config.Application, res http.ResponseWriter, err error) {
	stackTrace := fmt.Sprintf("[SERVER ERROR]: %s\n[STACK TRACE]:\n%s\n", err.Error(), debug.Stack())
	app.ErrorLog.Print(stackTrace)

	http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
