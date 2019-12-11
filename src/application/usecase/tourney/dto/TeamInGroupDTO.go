package dto

import (
	"application/usecase/player/dto"
	"domain/team/entity"
)

type TeamInGroupDTO struct {
	Player *dto.Player
	Team   *entity.Team //todo: convert entity to DTO
}

func (team TeamInGroupDTO) GetPlayer() *dto.Player {
	return team.Player
}

func NewTeamInGroupDTO(player *dto.Player, team *entity.Team) *TeamInGroupDTO {
	return &TeamInGroupDTO{Player: player, Team: team}
}
