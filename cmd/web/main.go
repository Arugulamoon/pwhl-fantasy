package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"

	"eden-walker.com/pwhl-fantasy/cmd/config"
	"eden-walker.com/pwhl-fantasy/pkg/fantasyteamsvc"
	"eden-walker.com/pwhl-fantasy/pkg/models"
	"eden-walker.com/pwhl-fantasy/pkg/models/filesystem"
	"eden-walker.com/pwhl-fantasy/pkg/playersvc"
)

func main() {
	fmt.Println("PWHL")

	ps := playersvc.PlayerService{}
	p := ps.Players()
	fmt.Printf("Players: %v\n", p)

	fts := fantasyteamsvc.FantasyTeamService{}

	ft := fts.GetFantasyTeams()
	fmt.Printf("Fantasy Teams: %v\n", ft)

	// LOGGING
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	// SESSION MANAGER
	sessionManager := scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true

	// HTML TEMPLATE CACHE
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// APPLICATION
	app := &config.Application{
		InfoLog:        infoLog,
		ErrorLog:       errorLog,
		SessionManager: sessionManager,
		TemplateCache:  templateCache,
		FantasyTeams: &filesystem.FantasyTeamModel{
			FantasyTeams: make([]*models.FantasyTeam, 0),
		},
		Players: &filesystem.PlayerModel{
			Players: make([]*models.Player, 0),
		},
	}

	// LOAD DATA
	app.Players.Load()

	// WEB SERVER
	srv := &http.Server{
		Addr:         ":9798",
		Handler:      Routes(app)(),
		ErrorLog:     errorLog,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting web server on %s...", srv.Addr)
	errorLog.Fatal(srv.ListenAndServe())
}
