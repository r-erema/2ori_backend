package dto

import (
	"toury_bakcend/src/application/usecase/player/dto"
	"toury_bakcend/src/domain/team/entity"
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
