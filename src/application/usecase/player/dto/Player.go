package dto

type Player struct {
	Name             string `json:"name"`
	TeamsCount       uint   `json:"teams_count"`
	RequiredTeamsIds []uint `json:"required_teams_ids"`
}
