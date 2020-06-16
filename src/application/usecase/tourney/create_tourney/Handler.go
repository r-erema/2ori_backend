package create_tourney

import (
	"toury_bakcend/src/application/service"
	"toury_bakcend/src/application/usecase/tourney/dto"
)

type Handler struct {
	teamsFiller service.TeamsFiller
}

func NewHandler(teamsFiller *service.TeamsFiller) *Handler {
	return &Handler{*teamsFiller}
}

func (handler Handler) Handle(command *Command) *dto.TourneyDTO {

	playerTeamBuckets := handler.teamsFiller.CreateAndFillBucketsForPlayers(command.players)

	var groupNamesMap = map[uint]string{
		0: "A",
		1: "B",
		2: "C",
		3: "D",
		4: "E",
		5: "F",
		6: "G",
		7: "H",
	}

	groupsCount := command.GetTeamsCount() / uint(4)
	var groups []*dto.GroupDTO
	for i := uint(0); i < groupsCount; i++ {
		groups = append(groups, dto.NewGroup(groupNamesMap[i], 4))
	}

	i := uint(0)
	getNextNotFullFilledGroup := func() (result *dto.GroupDTO) {
		allGroupsFullFilled := true

		for j := uint(0); j < groupsCount; j++ {
			group := groups[j]
			if !group.IsFullFilled() {
				allGroupsFullFilled = false
				break
			}
		}

		if !allGroupsFullFilled {
			group := groups[i]
			if !group.IsFullFilled() {
				result = group
			}
			if i == groupsCount-1 {
				i = 0
			} else {
				i++
			}
		}
		return result
	}

	for _, bucket := range playerTeamBuckets {
		for !bucket.IsEmpty() {
			group := getNextNotFullFilledGroup()
			if group == nil {
				break
			}
			group.AddTeam(dto.NewTeamInGroupDTO(bucket.Player(), bucket.PullOutRandomTeam()))
		}
	}

	for _, group := range groups {
		group.ShuffleTeams()
	}

	return dto.NewTourneyDTO(groups)
}
