package create_tourney

import "toury_bakcend/src/application/usecase/player/dto"

type Command struct {
	teamsCount uint
	players    []*dto.Player
}

func NewCommand(teamsCount uint, players []*dto.Player) *Command {
	return &Command{teamsCount: teamsCount, players: players}
}

func (command Command) GetTeamsCount() uint {
	return command.teamsCount
}

func (command Command) GetPlayers() []*dto.Player {
	return command.players
}
