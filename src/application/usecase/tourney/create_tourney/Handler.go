package create_tourney

import (
	"application/usecase/tourney/dto"
	"domain/team/entity"
	"domain/team/repository"
	"fmt"
	"math/rand"
	"time"
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
	var fetchedTeamsIds []string

	getFetchedTeamsIds := func(teams []entity.Team) []string {
		var ids []string
		for _, team := range teams {
			ids = append(ids, team.Id)
		}
		return ids
	}

	for _, player := range players {
		requiredTeams := handler.teamRepository.FindByIds(player.RequiredTeamsIds)
		fetchedTeamsIds = append(fetchedTeamsIds, getFetchedTeamsIds(requiredTeams)...)
		playerTeamsBuckets = append(playerTeamsBuckets, dto.NewBucket(player, requiredTeams))
	}
	otherTeams := handler.teamRepository.GetOrderedByRatingExceptIds(fetchedTeamsIds)

	currentPlayerIndex := 0
	otherTeamsCount := len(otherTeams)
	bucketsCount := len(playerTeamsBuckets)

	//todo: distribute teams per players
	for i := 0; i < otherTeamsCount; i++ {

		if currentPlayerIndex == bucketsCount {
			currentPlayerIndex = 0
		}

		if currentPlayerIndex < bucketsCount {
			bucket := playerTeamsBuckets[currentPlayerIndex]
			bucket.AppendTeams([]entity.Team{otherTeams[i]})
		}
		currentPlayerIndex++
	}

	fmt.Println(*playerTeamsBuckets[0], &playerTeamsBuckets[1])
}

func groupTeamByRating(teams []entity.Team) map[float32][]entity.Team {
	groupedByRating := map[float32][]entity.Team{}
	for _, team := range teams {
		groupedByRating[team.Rating] = append(groupedByRating[team.Rating], team)
	}
	return groupedByRating
}

func shuffle(teams []entity.Team) []entity.Team {
	var result []entity.Team
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(teams), func(i, j int) { teams[i], teams[j] = teams[j], teams[i] })
	return append(result, teams...)
}
