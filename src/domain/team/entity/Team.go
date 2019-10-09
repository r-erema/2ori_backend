package entity

type Team struct {
	id     string
	name   string
	rating float32
	league string
}

func (team *Team) GetId() string {
	return team.id
}