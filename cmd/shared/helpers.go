package shared

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"eden-walker.com/pwhl-fantasy/cmd/config"
)

func ServerError(app *config.Application) func(
	w http.ResponseWriter,
	err error,
) {
	return func(w http.ResponseWriter, err error) {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		app.ErrorLog.Output(2, trace)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func ClientError(app *config.Application) func(
	w http.ResponseWriter,
	status int,
) {
	return func(w http.ResponseWriter, status int) {
		app.InfoLog.Printf("%d %s", status, http.StatusText(status))
		http.Error(w, http.StatusText(status), status)
	}
}

func NotFound(app *config.Application) func(
	w http.ResponseWriter,
) {
	return func(w http.ResponseWriter) {
		ClientError(app)(w, http.StatusNotFound)
	}
}
