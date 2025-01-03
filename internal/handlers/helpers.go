package handlers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/JalalA984/apptrak/pkg/config"
)

func serverError(app *config.Application) func(res http.ResponseWriter, err error) {
	return func(res http.ResponseWriter, err error) {
		stackTrace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		app.ErrorLog.Print(stackTrace)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}
