package entity

type Tourney struct {
	Id         string `gorm:"primary_key"`
	TeamsCount uint   `gorm:"type:integer"`
}
