package dto

type Player struct {
	Name             string   `json:"name"`
	TeamsCount       uint     `json:"teamsCount"`
	RequiredTeamsIds []string `json:"requiredTeamsIds"`
}
