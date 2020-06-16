package dto

import "toury_bakcend/src/application/usecase/player/dto"

type TourneySettings struct {
	TourneyTeamsCount uint          `json:"tourneyTeamsCount"`
	Players           []*dto.Player `json:"players"`
}
