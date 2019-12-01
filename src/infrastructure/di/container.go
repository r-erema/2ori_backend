package di

import (
	"application/service"
	"application/usecase/tourney/create_tourney"
	team "domain/team/entity"
	"domain/team/repository"
	tourney "domain/tourney/entity"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv"
	"go.uber.org/dig"
	infrastructureRepo "infrastructure/repository"
	"log"
	"os"
)

func BuildContainer() *dig.Container {

	container := dig.New()

	err := container.Provide(func() *gorm.DB {
		dbUri := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
		db, err := gorm.Open(os.Getenv("DB_DRIVER"), dbUri)
		if err != nil {
			log.Fatal(err)
		}

		db.AutoMigrate(&team.Team{}, &tourney.Tourney{})

		return db
	})
	if err != nil {
		log.Fatal(err)
	}

	err = container.Provide(func(db *gorm.DB) repository.TeamRepositoryInterface {
		return infrastructureRepo.NewGormRepo(db)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = container.Provide(func(teamRepo repository.TeamRepositoryInterface) *service.TeamsFiller {
		return service.NewTeamsFiller(&teamRepo)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = container.Provide(func(teamsFiller *service.TeamsFiller) *create_tourney.Handler {
		return create_tourney.NewHandler(teamsFiller)
	})
	if err != nil {
		log.Fatal(err)
	}

	return container
}
