package create_tourney

import (
	"application/service"
	"application/usecase/tourney/dto"
)

type Handler struct {
	teamsFiller service.TeamsFiller
}

func NewHandler(teamsFiller *service.TeamsFiller) *Handler {
	return &Handler{*teamsFiller}
}

func (handler Handler) Handle(command *Command) {

	playerTeamBuckets := handler.teamsFiller.CreateAndFillBucketsForPlayers(command.players)

	groupsCount := command.GetTeamsCount() / uint(4)
	var groups []*dto.Group
	for i := uint(1); i <= groupsCount; i++ {
		groups = append(groups, dto.NewGroup(string(i)))
	}

	for _, group := range groups {

		for _, bucket := range playerTeamBuckets {
			group.AddTeam(dto.NewTeamInGroupDTO(bucket.Player(), bucket.PullOutRandomTeam()))
		}

	}

}
