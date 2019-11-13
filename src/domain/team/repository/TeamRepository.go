package repository

import "domain/team/entity"

type TeamRepositoryInterface interface {
	FindByIds(ids []uint) []entity.Team
}
