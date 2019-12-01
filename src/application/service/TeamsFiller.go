package service

import (
	playerDTO "application/usecase/player/dto"
	tourneyDTO "application/usecase/tourney/dto"
	"domain/team/entity"
	"domain/team/repository"
	"github.com/thoas/go-funk"
	"math/rand"
	"sort"
	"time"
)

type TeamsFiller struct {
	teamRepository repository.TeamRepositoryInterface
}

func NewTeamsFiller(teamRepository *repository.TeamRepositoryInterface) *TeamsFiller {
	return &TeamsFiller{teamRepository: *teamRepository}
}

func (filler TeamsFiller) CreateAndFillBucketsForPlayers(players []*playerDTO.Player) []*tourneyDTO.PlayerTeamsBucketDTO {
	buckets := filler.fillPlayersBucketsWithRequiredTeams(players)
	filler.fillBucketsByOtherTeams(buckets)
	return buckets
}

func (filler TeamsFiller) fillPlayersBucketsWithRequiredTeams(players []*playerDTO.Player) (playerTeamsBuckets []*tourneyDTO.PlayerTeamsBucketDTO) {
	for _, player := range players {
		requiredTeams := filler.teamRepository.FindByIds(player.RequiredTeamsIds)
		playerTeamsBuckets = append(playerTeamsBuckets, tourneyDTO.NewBucket(player, requiredTeams))
	}
	return
}

func (filler TeamsFiller) fillBucketsByOtherTeams(playerTeamsBuckets []*tourneyDTO.PlayerTeamsBucketDTO) {

	var fetchedTeamsIds []string
	funk.Map(playerTeamsBuckets, func(bucket *tourneyDTO.PlayerTeamsBucketDTO) interface{} {
		fetchedTeamsIds = append(fetchedTeamsIds, bucket.GetTeamsIds()...)
		return nil
	})
	otherTeams := filler.teamRepository.GetOrderedByRatingExceptIds(fetchedTeamsIds)
	otherTeams = shuffleTeamsByRatingGroup(otherTeams)

	otherTeamsCount := len(otherTeams)
	bucketsCount := len(playerTeamsBuckets)

	var ignoredIndexes []int
	startFromIndex := 0
	getNextNotFullFilledPlayerBucket := func() *tourneyDTO.PlayerTeamsBucketDTO {
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
}

func shuffleTeamsByRatingGroup(teams []entity.Team) []entity.Team {
	grouped := groupTeamByRating(teams)
	var result []entity.Team
	for _, group := range grouped {
		result = append(result, shuffle(group)...)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Rating > result[j].Rating
	})
	return result
}

func groupTeamByRating(teams []entity.Team) map[float32][]entity.Team {
	groupedByRating := map[float32][]entity.Team{}
	for _, team := range teams {
		groupedByRating[team.Rating] = append(groupedByRating[team.Rating], team)
	}
	return groupedByRating
}

//todo: replace with shuffle in go-funk package
func shuffle(teams []entity.Team) []entity.Team {
	var result []entity.Team
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(teams), func(i, j int) { teams[i], teams[j] = teams[j], teams[i] })
	return append(result, teams...)
}
