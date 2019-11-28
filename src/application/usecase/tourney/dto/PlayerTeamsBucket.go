package dto

import (
	"application/usecase/player/dto"
	"domain/team/entity"
	"github.com/thoas/go-funk"
)

type PlayerTeamsBucket struct {
	player *dto.Player
	teams  []entity.Team
}

func (bucket *PlayerTeamsBucket) Player() *dto.Player {
	return bucket.player
}

func NewBucket(player *dto.Player, teams []entity.Team) *PlayerTeamsBucket {
	return &PlayerTeamsBucket{
		player: player,
		teams:  teams,
	}
}

func (bucket *PlayerTeamsBucket) AppendTeams(teams []entity.Team) {
	bucket.teams = append(bucket.teams, teams...)
}

func (bucket *PlayerTeamsBucket) TeamsCount() uint {
	return uint(len(bucket.teams))
}

func (bucket *PlayerTeamsBucket) GetTeamsIds() []string {
	return funk.Map(bucket.teams, func(team entity.Team) string { return team.Id }).([]string)
}
