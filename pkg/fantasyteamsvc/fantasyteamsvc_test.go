package fantasyteamsvc_test

import (
	"testing"

	"eden-walker.com/pwhl-fantasy/pkg/fantasyteamsvc"
)

func TestInitialization(t *testing.T) {
	fts := &fantasyteamsvc.FantasyTeamService{
		FantasyTeams: make([]string, 0),
	}

	got := fts.GetFantasyTeams()
	want := []string{}

	if len(got) != len(want) {
		t.Errorf("want %d; got %d", len(want), len(got))
	}
	for k, v := range got {
		if v != want[k] {
			t.Errorf("want %s; got %s", want[k], v)
		}
	}
}

func TestAddFantasyTeam(t *testing.T) {
	fts := &fantasyteamsvc.FantasyTeamService{
		FantasyTeams: make([]string, 0),
	}

	got := fts.AddFantasyTeam("Nicholas' Team")
	want := "Nicholas' Team"

	if got != want {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestGetFantasyTeams(t *testing.T) {
	fts := &fantasyteamsvc.FantasyTeamService{
		FantasyTeams: make([]string, 0),
	}
	fts.AddFantasyTeam("Nicholas' Team")

	got := fts.GetFantasyTeams()
	want := []string{"Nicholas' Team"}

	if len(got) != len(want) {
		t.Errorf("want %d; got %d", len(want), len(got))
	}
	for k, v := range got {
		if v != want[k] {
			t.Errorf("want %s; got %s", want[k], v)
		}
	}
}

func TestGetFantasyTeamsMultiple(t *testing.T) {
	fts := &fantasyteamsvc.FantasyTeamService{
		FantasyTeams: make([]string, 0),
	}
	fts.AddFantasyTeam("Nicholas' Team")
	fts.AddFantasyTeam("Deirdre's Team")

	got := fts.GetFantasyTeams()
	want := []string{"Nicholas' Team", "Deirdre's Team"}

	if len(got) != len(want) {
		t.Errorf("want %d; got %d", len(want), len(got))
	}
	for k, v := range got {
		if v != want[k] {
			t.Errorf("want %s; got %s", want[k], v)
		}
	}
}
