package repository

import (
	"github.com/jinzhu/gorm"
	"toury_bakcend/src/domain/team/entity"
	repository "toury_bakcend/src/domain/team/repository"
)

type GormTeamsRepo struct {
	db *gorm.DB
}

func NewGormRepo(db *gorm.DB) repository.TeamRepositoryInterface {
	return GormTeamsRepo{db: db}
}

func (Repo GormTeamsRepo) GetAll() []*entity.Team {
	var teams []*entity.Team
	Repo.db.Find(&teams)
	return teams
}

func (Repo GormTeamsRepo) FindByIds(ids []string) []entity.Team {
	var teams []entity.Team
	Repo.db.Where(ids).Find(&teams)
	return teams
}

func (Repo GormTeamsRepo) GetOrderedByRatingExceptIds(exceptIds []string) []entity.Team {
	var teams []entity.Team
	Repo.db.Not(exceptIds).Order("rating desc").Find(&teams)
	return teams
}
