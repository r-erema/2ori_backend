package dto

import (
	"application/usecase/player/dto"
	"domain/team/entity"
)

type TeamInGroupDTO struct {
	player *dto.Player
	team   *entity.Team //todo: convert entity to DTO
}

func NewTeamInGroupDTO(player *dto.Player, team *entity.Team) *TeamInGroupDTO {
	return &TeamInGroupDTO{player: player, team: team}
}
