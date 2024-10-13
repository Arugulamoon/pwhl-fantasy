package playersvc

type PlayerService struct{}

func (ps *PlayerService) Players() []string {
	return []string{"Brianne Jenner", "Emily Clark"}
}
