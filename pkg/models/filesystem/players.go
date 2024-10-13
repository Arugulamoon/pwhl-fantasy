package filesystem

import "eden-walker.com/pwhl-fantasy/pkg/models"

type PlayerModel struct {
	Players []*models.Player
}

func (m *PlayerModel) Load() {
	m.Players = append(m.Players, &models.Player{Id: 1, Name: "Brianne Jenner"})
	m.Players = append(m.Players, &models.Player{Id: 2, Name: "Emily Clark"})
}

func (m *PlayerModel) GetAll() ([]*models.Player, error) {
	return m.Players, nil
}

func (m *PlayerModel) Get(id int) (*models.Player, error) {
	if id > len(m.Players) {
		return nil, models.ErrNoRecord
	}
	return m.Players[id-1], nil
}
