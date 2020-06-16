package dto

import (
	"toury_bakcend/src/domain/team/entity"
)

type TeamsDTO struct {
	Teams []*entity.Team
}

func NewTeamsDTO(teams []*entity.Team) *TeamsDTO {
	return &TeamsDTO{Teams: teams}
}
