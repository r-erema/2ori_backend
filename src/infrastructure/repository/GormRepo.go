package repository

import (
	"domain/team/entity"
	"github.com/jinzhu/gorm"
)

type GormTeamsRepo struct {
	db *gorm.DB
}

func NewGormRepo(db *gorm.DB) *GormTeamsRepo {
	return &GormTeamsRepo{db: db}
}

func (Repo *GormTeamsRepo) FindByIds(ids []string) []entity.Team {
	var teams []entity.Team
	Repo.db.Where(ids).Find(&teams)
	return teams
}

func (Repo *GormTeamsRepo) GetOrderedByRatingExceptIds(exceptIds []string) []entity.Team {
	var teams []entity.Team
	Repo.db.Not(exceptIds).Order("rating desc").Find(&teams)
	return teams
}
