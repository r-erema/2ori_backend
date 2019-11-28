package dto

type Group struct {
	name  string
	teams []*TeamInGroupDTO
}

func NewGroup(name string) *Group {
	return &Group{name: name}
}

func (group Group) AddTeam(team *TeamInGroupDTO) {
	group.teams = append(group.teams, team)
}
