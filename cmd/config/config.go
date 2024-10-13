package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"

	"eden-walker.com/pwhl-fantasy/pkg/models/filesystem"
)

type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger

	SessionManager *scs.SessionManager
	TemplateCache  map[string]*template.Template

	FantasyTeams *filesystem.FantasyTeamModel
	Players      *filesystem.PlayerModel
}
