package shared

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"

	"eden-walker.com/pwhl-fantasy/cmd/config"
)

func LogRequest(app *config.Application) func(
	next http.Handler,
) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			app.InfoLog.Printf("%s - %s %s %s",
				r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

			next.ServeHTTP(w, r)
		})
	}
}

func RecoverPanic(app *config.Application) func(
	next http.Handler,
) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.Header().Set("Connection", "close")
					ServerError(app)(w, fmt.Errorf("%s", err))
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})

	return csrfHandler
}
