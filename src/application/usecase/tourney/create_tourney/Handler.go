package create_tourney

import (
	"../dto"
	"domain/team/repository"
)

type Handler struct {
	teamRepository *repository.TeamRepository
}

func NewHandler(teamRepository *repository.TeamRepository) *Handler {
	return &Handler{teamRepository}
}

func (handler Handler) Handle(command Command) {

	players := command.getPlayers()
	var playerTeamsBuckets []*dto.PlayerTeamsBucket
	channel := make(chan *dto.PlayerTeamsBucket)
	for _, player := range players {
		go func() {
			channel <- dto.NewBucket(player, handler.teamRepository.FindByIds(player.GetRequiredTeamsIds()))
		}()
	}

	select {
	case bucket := <-channel:
		playerTeamsBuckets = append(playerTeamsBuckets, bucket)
	}

}
