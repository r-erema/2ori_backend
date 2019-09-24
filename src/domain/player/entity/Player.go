package entity

import (
	"domain/team/entity"
)

type Player struct {
	name          string
	teamsCount    uint
	requiredTeams []entity.Team
}
