package create_tourney_test

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"gopkg.in/testfixtures.v2"
	"log"
	"testing"
	"toury_bakcend/src/application/usecase/player/dto"
	"toury_bakcend/src/application/usecase/tourney/create_tourney"
	TourneyDTO "toury_bakcend/src/application/usecase/tourney/dto"
	"toury_bakcend/src/infrastructure/di"
)

var (
	container = di.BuildContainer()
)

func TestHandle(t *testing.T) {
	loadDotEnv()
	prepareTestDatabase()

	var createTourneyUseCase = create_tourney.NewCommand(8, []*dto.Player{
		{
			Name:            "Player 1",
			TeamsCount:      2,
			RequiredTeamIds: []string{"58b1d3c5-d993-44ac-86cb-e1771ed01f5e"},
		},
		{
			Name:            "Player 2",
			TeamsCount:      4,
			RequiredTeamIds: []string{"8f5a84ae-d239-4933-8376-08d887e85404", "bac5c0f8-7d98-4097-b091-6d251ebf83ad"},
		},
		{
			Name:            "Player 3",
			TeamsCount:      2,
			RequiredTeamIds: []string{},
		},
	})

	var tourney *TourneyDTO.TourneyDTO
	err := container.Invoke(func(handler *create_tourney.Handler) {
		tourney = handler.Handle(createTourneyUseCase)
	})
	if err != nil {
		t.Error(err)
	}

	if len(tourney.GetGroups()) != 2 {
		t.Error()
	}

}

func loadDotEnv() {
	if err := godotenv.Load("../../../../../.env.test"); err != nil {
		log.Fatal(err)
	}
}

func prepareTestDatabase() {
	_ = container.Invoke(func(db *gorm.DB) {
		fixtures, err := testfixtures.NewFolder(db.DB(), &testfixtures.PostgreSQL{}, "../../../../../fixture")
		if err != nil {
			log.Fatal(err)
		}
		if err := fixtures.Load(); err != nil {
			log.Fatal(err)
		}
	})
}
