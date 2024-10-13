package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/justinas/nosurf"

	"eden-walker.com/pwhl-fantasy/cmd/config"
	"eden-walker.com/pwhl-fantasy/cmd/shared"
)

func AddDefaultData(app *config.Application) func(
	td *templateData,
	r *http.Request,
) *templateData {
	return func(td *templateData, r *http.Request) *templateData {
		if td == nil {
			td = &templateData{}
		}
		td.CSRFToken = nosurf.Token(r)
		td.Flash = app.SessionManager.PopString(r.Context(), "flash")
		td.CurrentYear = time.Now().Year()
		return td
	}
}

func Render(app *config.Application) func(
	w http.ResponseWriter,
	r *http.Request,
	name string,
	td *templateData,
) {
	return func(
		w http.ResponseWriter,
		r *http.Request,
		name string,
		td *templateData,
	) {
		ts, ok := app.TemplateCache[name]
		if !ok {
			shared.ServerError(app)(w,
				fmt.Errorf("template %s does not exist", name))
			return
		}

		buf := new(bytes.Buffer)
		err := ts.Execute(buf, AddDefaultData(app)(td, r))
		if err != nil {
			shared.ServerError(app)(w, err)
			return
		}

		buf.WriteTo(w)
	}
}
