package repository

import (
	"domain/team/entity"
	repository "domain/team/repository"
	"github.com/jinzhu/gorm"
)

type GormTeamsRepo struct {
	db *gorm.DB
}

func NewGormRepo(db *gorm.DB) repository.TeamRepositoryInterface {
	return GormTeamsRepo{db: db}
}

func (Repo GormTeamsRepo) GetAll() []*entity.Team {
	var teams []*entity.Team
	Repo.db.Find(teams)
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
