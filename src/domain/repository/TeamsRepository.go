package repository

import "domain/entity"

type TeamsRepository interface {
	Find(ID int64) (*entity.Team, error)
}
