package dto

type Player struct {
	name             string
	teamsCount       uint
	requiredTeamsIds []uint
}

func (player Player) GetRequiredTeamsIds() []uint {
	return player.requiredTeamsIds
}
