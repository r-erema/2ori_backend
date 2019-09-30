package dto

type Player struct {
	name             string
	teamsCount       uint
	requiredTeamsIds []uint8
}

func (player Player) GetRequiredTeamsIds() []uint8 {
	return player.requiredTeamsIds
}
