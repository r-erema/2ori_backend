package get_teams

import (
	"application/usecase/team/dto"
	"domain/team/repository"
)

type Handler struct {
	teamRepository repository.TeamRepositoryInterface
}

func NewHandler(teamRepository *repository.TeamRepositoryInterface) *Handler {
	return &Handler{teamRepository: *teamRepository}
}

func (handler Handler) Handle() *dto.TeamsDTO {
	return dto.NewTeamsDTO(handler.teamRepository.GetAll())
}
