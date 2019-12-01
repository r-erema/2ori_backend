package dto

import "application/usecase/player/dto"

type TourneySettings struct {
	TourneyTeamsCount uint `json:"tourney_teams_count"`
	Players           []*dto.Player
}
