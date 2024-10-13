package filesystem

import "eden-walker.com/pwhl-fantasy/pkg/models"

type FantasyTeamModel struct {
	FantasyTeams []*models.FantasyTeam
}

func (m *FantasyTeamModel) GetAll() ([]*models.FantasyTeam, error) {
	return m.FantasyTeams, nil
}

func (m *FantasyTeamModel) Insert(name string) (int, error) {
	id := len(m.FantasyTeams) + 1
	m.FantasyTeams = append(m.FantasyTeams, &models.FantasyTeam{
		Id:   id,
		Name: name,
	})
	return id, nil
}

func (m *FantasyTeamModel) Get(id int) (*models.FantasyTeam, error) {
	if id > len(m.FantasyTeams) {
		return nil, models.ErrNoRecord
	}
	return m.FantasyTeams[id-1], nil
}
