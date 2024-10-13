package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"eden-walker.com/pwhl-fantasy/cmd/config"
	"eden-walker.com/pwhl-fantasy/cmd/shared"
	"eden-walker.com/pwhl-fantasy/pkg/forms"
	"eden-walker.com/pwhl-fantasy/pkg/models"
)

// Fantasy Teams
func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fantasyTeams, err := app.FantasyTeams.GetAll()
		if err != nil {
			shared.ServerError(app)(w, err)
			return
		}

		Render(app)(w, r, "home.page.tmpl", &templateData{
			FantasyTeams: fantasyTeams,
		})
	}
}

func ShowFantasyTeam(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil || id < 1 {
			shared.NotFound(app)(w)
			return
		}

		fantasyTeam, err := app.FantasyTeams.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				shared.NotFound(app)(w)
			} else {
				shared.ServerError(app)(w, err)
			}
			return
		}

		Render(app)(w, r, "show.page.tmpl", &templateData{
			FantasyTeam: fantasyTeam,
		})
	}
}

func NewFantasyTeamForm(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Render(app)(w, r, "create.page.tmpl", &templateData{
			Form: forms.New(nil),
		})
	}
}

func CreateFantasyTeam(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			shared.ClientError(app)(w, http.StatusBadRequest)
			return
		}

		form := forms.New(r.PostForm)
		form.Required("name")
		form.MaxLength("name", 100)

		if !form.Valid() {
			Render(app)(w, r, "create.page.tmpl", &templateData{
				Form: form,
			})
			return
		}

		id, err := app.FantasyTeams.Insert(form.Get("name"))
		if err != nil {
			shared.ServerError(app)(w, err)
			return
		}

		app.SessionManager.Put(r.Context(), "flash",
			"Fantasy Team created successfully!")

		http.Redirect(w, r,
			fmt.Sprintf("/fantasyteams/%d", id), http.StatusSeeOther)
	}
}

// Players
func ListAllPlayers(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		players, err := app.Players.GetAll()
		if err != nil {
			shared.ServerError(app)(w, err)
			return
		}

		Render(app)(w, r, "players.page.tmpl", &templateData{
			Players: players,
		})
	}
}

func ShowPlayer(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil || id < 1 {
			shared.NotFound(app)(w)
			return
		}

		player, err := app.Players.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				shared.NotFound(app)(w)
			} else {
				shared.ServerError(app)(w, err)
			}
			return
		}

		Render(app)(w, r, "player.page.tmpl", &templateData{
			Player: player,
		})
	}
}
