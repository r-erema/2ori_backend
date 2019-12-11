package dto

import "github.com/thoas/go-funk"

type GroupDTO struct {
	Name       string
	Teams      []*TeamInGroupDTO
	TeamsCount uint
}

func (group *GroupDTO) GetTeams() []*TeamInGroupDTO {
	return group.Teams
}

func (group *GroupDTO) ShuffleTeams() {
	group.Teams = funk.Shuffle(group.Teams).([]*TeamInGroupDTO)
}

func NewGroup(name string, teamsCount uint) *GroupDTO {
	return &GroupDTO{Name: name, TeamsCount: teamsCount}
}

func (group *GroupDTO) AddTeam(team *TeamInGroupDTO) {
	group.Teams = append(group.Teams, team)
}

func (group *GroupDTO) IsFullFilled() bool {
	return uint(len(group.Teams)) == group.TeamsCount
}
