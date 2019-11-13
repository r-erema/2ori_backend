package create_tourney

import "application/usecase/player/dto"

type Command struct {
	teamsCount uint
	players    []*dto.Player
}

func NewCommand(teamsCount uint, players []*dto.Player) *Command {
	return &Command{teamsCount: teamsCount, players: players}
}

func (command Command) getTeamsCount() uint {
	return command.teamsCount
}

func (command Command) getPlayers() []*dto.Player {
	return command.players
}
