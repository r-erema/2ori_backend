package dto

type Player struct {
	Name            string   `json:"name"`
	TeamsCount      uint     `json:"teamsCount"`
	RequiredTeamIds []string `json:"requiredTeamIds"`
}
