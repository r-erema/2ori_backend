package create_tourney

import (
	"../dto"
	"domain/team/repository"
)

type handler struct {
	teamRepository repository.TeamRepository
}

func NewHandler(teamRepository *repository.TeamRepository) *handler {
	return &handler{teamRepository}
}

func (handler handler) Handle(command Command) {

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
