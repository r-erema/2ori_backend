package dto

type Player struct {
	name          string
	teamsCount    uint
	requiredTeams []*Team
}

func (player *Player) getName() string {
	return player.name
}
