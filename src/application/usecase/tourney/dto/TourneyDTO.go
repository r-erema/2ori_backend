package dto

type TourneyDTO struct {
	Groups []*GroupDTO
}

func NewTourneyDTO(groups []*GroupDTO) *TourneyDTO {
	return &TourneyDTO{Groups: groups}
}

func (tourney TourneyDTO) GetGroups() []*GroupDTO {
	return tourney.Groups
}
