package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record was found")

type FantasyTeam struct {
	Id   int
	Name string

	Players []*FantasyTeamPlayer
}

type FantasyTeamPlayer struct {
	Player     *Player
	Start, End time.Time
}

type Player struct {
	Id   int
	Name string

	FantasyTeam *FantasyTeam
}
