package entity

type Team struct {
	Id     string  `gorm:"primary_key"`
	Name   string  `gorm:"type:varchar(100)"`
	Rating float32 `gorm:"type:decimal(2,1)"`
	League string  `gorm:"type:varchar(50)"`
}
