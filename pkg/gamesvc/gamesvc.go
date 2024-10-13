package gamesvc

import "github.com/google/uuid"

type Goal struct {
	Scorer          string
	PrimaryAssist   string
	SecondaryAssist string
	Type            string
}

type Game struct {
	ID    uuid.UUID
	Goals []*Goal
}

func (gm *Game) AddGoal(scorer, assist1, assist2 string) *Goal {
	gl := &Goal{
		Scorer:          scorer,
		PrimaryAssist:   assist1,
		SecondaryAssist: assist2,
	}
	gm.Goals = append(gm.Goals, gl)
	return gl
}

type GameService struct {
	Games []*Game
}

func (gsvc *GameService) AddGame() *Game {
	gm := &Game{ID: uuid.New()}
	gsvc.Games = append(gsvc.Games, gm)
	return gm
}
