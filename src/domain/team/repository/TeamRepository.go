package repository

import "domain/team/entity"

type TeamRepository interface {
	FindByIds(ids []uint8) []entity.Team
}
