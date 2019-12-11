package dto

import "application/usecase/player/dto"

type TourneySettings struct {
	TourneyTeamsCount uint `json:"tourneyTeamsCount"`
	Players           []*dto.Player
}
