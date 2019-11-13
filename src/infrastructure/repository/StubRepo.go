package repository

import (
	"domain/team/entity"
)

type StubRepo struct {
}

func (Repo *StubRepo) FindByIds(ids []uint) []entity.Team {
	return []entity.Team{}
}
