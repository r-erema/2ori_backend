package dto

import (
	"../../player/dto"
	"domain/team/entity"
)

type PlayerTeamsBucket struct {
	player *dto.Player
	teams  []entity.Team
}

func NewBucket(player *dto.Player, teams []entity.Team) *PlayerTeamsBucket {
	return &PlayerTeamsBucket{
		player: player,
		teams:  teams,
	}
}
