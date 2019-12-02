package dto

import (
	"domain/team/entity"
)

type TeamsDTO struct {
	teams []*entity.Team
}

func NewTeamsDTO(teams []*entity.Team) *TeamsDTO {
	return &TeamsDTO{teams: teams}
}
