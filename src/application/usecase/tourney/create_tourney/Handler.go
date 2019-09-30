package create_tourney

import (
	"domain/team/entity"
	"domain/team/repository"
	"github.com/thoas/go-funk"
)

type Handler struct {
	teamRepository repository.TeamRepository
}

func (handler Handler) handle(command Command) {

	players := command.getPlayers()
	var ids []uint8
	for _, player := range players {
		ids = append(ids, player.GetRequiredTeamsIds()...)
	}

	teams := handler.teamRepository.FindByIds(ids)

	for _, player := range players {
		funk.Filter(teams, func(team entity.Team) bool {
			return funk.Contains(player.GetRequiredTeamsIds(), team.GetId())
		})
	}

}
