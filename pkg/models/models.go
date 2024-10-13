package models

import "errors"

var ErrNoRecord = errors.New("models: no matching record was found")

type FantasyTeam struct {
	Id   int
	Name string
}

type Player struct {
	Id   int
	Name string
}
