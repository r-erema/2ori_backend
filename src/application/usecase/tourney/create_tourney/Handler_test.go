package create_tourney_test

import (
	"application/usecase/player/dto"
	"application/usecase/tourney/create_tourney"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"gopkg.in/testfixtures.v2"
	"infrastructure/di"
	"log"
	"testing"
)

var (
	container = di.BuildContainer()
)

func TestHandle(t *testing.T) {
	loadDotEnv()
	prepareTestDatabase()

	var createTourneyUseCase = create_tourney.NewCommand(16, []*dto.Player{
		{
			Name:             "Player 1",
			TeamsCount:       2,
			RequiredTeamsIds: []string{"58b1d3c5-d993-44ac-86cb-e1771ed01f5e"},
		},
		{
			Name:             "Player 2",
			TeamsCount:       4,
			RequiredTeamsIds: []string{"8f5a84ae-d239-4933-8376-08d887e85404", "bac5c0f8-7d98-4097-b091-6d251ebf83ad"},
		},
		{
			Name:             "Player 2",
			TeamsCount:       2,
			RequiredTeamsIds: []string{},
		},
	})

	err := container.Invoke(func(createTourneyHandler *create_tourney.Handler) {
		createTourneyHandler.Handle(createTourneyUseCase)
	})

	if err != nil {
		t.Error(err)
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
