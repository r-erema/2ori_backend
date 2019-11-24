package repository

import "domain/team/entity"

type TeamRepositoryInterface interface {
	FindByIds(ids []string) []entity.Team
	GetOrderedByRatingExceptIds(exceptIds []string) []entity.Team
}
