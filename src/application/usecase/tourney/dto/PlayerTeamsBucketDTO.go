package dto

import (
	"application/usecase/player/dto"
	"domain/team/entity"
	"github.com/thoas/go-funk"
	"math/rand"
	"time"
)

type PlayerTeamsBucketDTO struct {
	player *dto.Player

	//todo: convert entity to DTO
	teams []entity.Team
}

func (bucket *PlayerTeamsBucketDTO) Player() *dto.Player {
	return bucket.player
}

func NewBucket(player *dto.Player, teams []entity.Team) *PlayerTeamsBucketDTO {
	return &PlayerTeamsBucketDTO{
		player: player,
		teams:  teams,
	}
}

func (bucket *PlayerTeamsBucketDTO) AppendTeams(teams []entity.Team) {
	bucket.teams = append(bucket.teams, teams...)
}

func (bucket *PlayerTeamsBucketDTO) TeamsCount() uint {
	return uint(len(bucket.teams))
}

func (bucket *PlayerTeamsBucketDTO) GetTeamsIds() []string {
	return funk.Map(bucket.teams, func(team entity.Team) string { return team.Id }).([]string)
}

func (bucket *PlayerTeamsBucketDTO) PullOutRandomTeam() *entity.Team {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(bucket.teams))
	chosen := bucket.teams[i]
	bucket.teams = append(bucket.teams[:i], bucket.teams[i+1:]...)
	return &chosen
}

func (bucket *PlayerTeamsBucketDTO) IsEmpty() bool {
	return len(bucket.teams) == 0
}
