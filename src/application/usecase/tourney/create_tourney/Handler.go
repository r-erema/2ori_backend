package create_tourney

import (
	"application/usecase/tourney/dto"
	"domain/team/entity"
	"domain/team/repository"
	"fmt"
	"github.com/thoas/go-funk"
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
	otherTeams = shuffleTeamsByRatingGroup(otherTeams)

	otherTeamsCount := len(otherTeams)
	bucketsCount := len(playerTeamsBuckets)

	var ignoredIndexes []int
	startFromIndex := 0
	getNextNotFullFilledPlayerBucket := func() *dto.PlayerTeamsBucket {
		for i := startFromIndex; i < bucketsCount; i++ {
			if funk.IndexOf(ignoredIndexes, i) == -1 {
				if startFromIndex >= bucketsCount-1 {
					startFromIndex = 0
				} else {
					startFromIndex = i + 1
				}

				bucket := playerTeamsBuckets[i]
				if bucket.TeamsCount() >= bucket.Player().TeamsCount {
					ignoredIndexes = append(ignoredIndexes, i)
					return nil
				}

				return bucket
			}
		}
		return nil
	}

	for i := 0; i < otherTeamsCount; i++ {
		bucket := getNextNotFullFilledPlayerBucket()
		if bucket != nil {
			bucket.AppendTeams([]entity.Team{otherTeams[i]})
		}
	}

	fmt.Println(*playerTeamsBuckets[0], &playerTeamsBuckets[1])
}

func shuffleTeamsByRatingGroup(teams []entity.Team) []entity.Team {
	grouped := groupTeamByRating(teams)
	var result []entity.Team
	for _, group := range grouped {
		result = append(result, shuffle(group)...)
	}
	return result
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
