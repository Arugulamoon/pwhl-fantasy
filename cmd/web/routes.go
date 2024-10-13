package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"eden-walker.com/pwhl-fantasy/cmd/config"
	"eden-walker.com/pwhl-fantasy/cmd/shared"
)

func Routes(app *config.Application) func() *mux.Router {
	return func() *mux.Router {
		// TODO: Configure router with not found handler in helpers.go
		r := mux.NewRouter()

		// Middleware
		r.Use(shared.RecoverPanic(app))
		r.Use(shared.LogRequest(app))
		r.Use(shared.NoSurf)
		r.Use(app.SessionManager.LoadAndSave)

		// Fantasy Teams Routes
		r.HandleFunc("/", Home(app)).
			Methods(http.MethodGet, http.MethodHead)
		r.HandleFunc("/fantasyteams/{id:[0-9]+}", ShowFantasyTeam(app)).
			Methods(http.MethodGet, http.MethodHead)
		r.HandleFunc("/fantasyteams/new", NewFantasyTeamForm(app)).
			Methods(http.MethodGet, http.MethodHead)
		r.HandleFunc("/fantasyteams", CreateFantasyTeam(app)).
			Methods(http.MethodPost)

		// Players Routes
		r.HandleFunc("/players", ListAllPlayers(app)).
			Methods(http.MethodGet, http.MethodHead)
		r.HandleFunc("/players/{id:[0-9]+}", ShowPlayer(app)).
			Methods(http.MethodGet, http.MethodHead)

		// Static Files
		r.HandleFunc("/static", http.NotFound)
		sf := r.PathPrefix("/static/").Subrouter()
		fileServer := http.FileServer(protectedFileSystem{http.Dir("./ui/static/")})
		sf.PathPrefix("/").Handler(http.StripPrefix("/static", fileServer))

		return r
	}
}
