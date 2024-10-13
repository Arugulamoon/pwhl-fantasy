package main

import (
	"html/template"
	"path/filepath"
	"time"

	"eden-walker.com/pwhl-fantasy/pkg/forms"
	"eden-walker.com/pwhl-fantasy/pkg/models"
)

type templateData struct {
	CurrentYear int
	Flash       string
	Form        *forms.Form
	CSRFToken   string

	FantasyTeams []*models.FantasyTeam
	FantasyTeam  *models.FantasyTeam

	Players []*models.Player
	Player  *models.Player
}

func humanReadableDate(t time.Time) string {
	// Magical Date for Formatting: 01/02 03:04:05PM '06 -0700
	return t.Format("Jan 02 2006 at 3:04:05 PM MST")
}

var functions = template.FuncMap{
	"humanReadableDate": humanReadableDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	// Get slice of all filepaths with extension '.page.tmpl'
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extract file name (like 'home.page.tmpl') from full file path
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add 'layout' templates to template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// Add 'partial' templates to template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		// Add template set to cache,
		// using name of page (eg 'home.page.tmpl') as key
		cache[name] = ts
	}

	// Return the map.
	return cache, nil
}
