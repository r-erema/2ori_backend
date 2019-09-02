package use_case

import "application/dto"

type CreateTourney struct {
	teamsCount uint
	players    []*dto.Player
}

func NewCreateTourney(teamsCount uint, players []*dto.Player) *CreateTourney {
	return &CreateTourney{teamsCount, players}
}
