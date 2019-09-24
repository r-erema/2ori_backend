package create_tourney

import "domain/team/repository"

type Handler struct {
	teamRepository repository.TeamRepository
}

func (handler Handler) handle(command Command) {

	players := command.getPlayers()
	var ids []uint
	for _, player := range players {
		ids = append(ids, player.GetRequiredTeamsIds()...)
	}

	//handler.teamRepository.FindByIds([])

}
