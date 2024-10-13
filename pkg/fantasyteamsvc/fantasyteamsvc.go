package fantasyteamsvc

type FantasyTeamService struct {
	FantasyTeams []string
}

func (fts *FantasyTeamService) GetFantasyTeams() []string {
	return fts.FantasyTeams
}

func (fts *FantasyTeamService) AddFantasyTeam(
	name string,
) string {
	fts.FantasyTeams = append(fts.FantasyTeams, name)
	return name
}
