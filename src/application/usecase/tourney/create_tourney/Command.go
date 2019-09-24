package create_tourney

import "../../player/dto"

type Command struct {
	teamsCount uint
	players    []*dto.Player
}

func (command Command) getTeamsCount() uint {
	return command.teamsCount
}

func (command Command) getPlayers() []*dto.Player {
	return command.players
}
