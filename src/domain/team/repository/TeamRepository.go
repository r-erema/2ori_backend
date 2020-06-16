package repository

import "toury_bakcend/src/domain/team/entity"

type TeamRepositoryInterface interface {
	GetAll() []*entity.Team
	FindByIds(ids []string) []entity.Team
	GetOrderedByRatingExceptIds(exceptIds []string) []entity.Team
}
