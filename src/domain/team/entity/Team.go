package entity

type Team struct {
	id     string  `gorm:"primary_key"`
	name   string  `gorm:"type:varchar(100);`
	rating float32 `gorm:"type:float(2,1);`
	league string  `gorm:"type:varchar(50);`
}

func (team *Team) GetId() string {
	return team.id
}
