package entity

type Tourney struct {
	id         string
	teamsCount uint
}

func NewTourney(id string, teamsCount uint) *Tourney {
	return &Tourney{id, teamsCount}
}

/*func (this *Tourney) runToss() dto.TossResult {

}*/
