package create_tourney

import (
	"application/usecase/tourney/dto"
	"domain/team/repository"
	"fmt"
	"sync"
)

type Handler struct {
	teamRepository repository.TeamRepositoryInterface
}

func NewHandler(teamRepository *repository.TeamRepositoryInterface) *Handler {
	return &Handler{*teamRepository}
}

func (handler Handler) Handle(command *Command) {

	players := command.getPlayers()
	var playerTeamsBuckets []*dto.PlayerTeamsBucket
	channel := make(chan *dto.PlayerTeamsBucket)

	wg := sync.WaitGroup{}
	for _, player := range players {
		wg.Add(1)
		go func() {
			channel <- dto.NewBucket(player, handler.teamRepository.FindByIds(player.RequiredTeamsIds))
			wg.Done()
		}()
	}

	select {
	case bucket := <-channel:
		playerTeamsBuckets = append(playerTeamsBuckets, bucket)
	}
	select {
	case bucket := <-channel:
		playerTeamsBuckets = append(playerTeamsBuckets, bucket)
	}

	wg.Wait()

	fmt.Println(*playerTeamsBuckets[0], &playerTeamsBuckets[1])

}
