package dto

type TourneyDTO struct {
	groups []*GroupDTO
}

func NewTourneyDTO(groups []*GroupDTO) *TourneyDTO {
	return &TourneyDTO{groups: groups}
}

func (tourney TourneyDTO) GetGroups() []*GroupDTO {
	return tourney.groups
}
