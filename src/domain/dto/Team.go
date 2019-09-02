package dto

type Team struct {
	name   string
	rating float64
	league League
	player Player
}

func NewTeam(name string, rating float64, league League, player Player) *Team {
	return &Team{name, rating, league, player}
}
