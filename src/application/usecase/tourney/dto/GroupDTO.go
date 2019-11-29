package dto

import "github.com/thoas/go-funk"

type GroupDTO struct {
	name       string
	teams      []*TeamInGroupDTO
	teamsCount uint
}

func (group *GroupDTO) GetTeams() []*TeamInGroupDTO {
	return group.teams
}

func (group *GroupDTO) ShuffleTeams() {
	group.teams = funk.Shuffle(group.teams).([]*TeamInGroupDTO)
}

func NewGroup(name string, teamsCount uint) *GroupDTO {
	return &GroupDTO{name: name, teamsCount: teamsCount}
}

func (group *GroupDTO) AddTeam(team *TeamInGroupDTO) {
	group.teams = append(group.teams, team)
}

func (group *GroupDTO) IsFullFilled() bool {
	return uint(len(group.teams)) == group.teamsCount
}
