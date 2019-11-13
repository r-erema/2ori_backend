package main

import (
	"application/usecase/player/dto"
	"application/usecase/tourney/create_tourney"
	"config"
	"domain/team/repository"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"go.uber.org/dig"
	infrastructureRepo "infrastructure/repository"
	"log"
	"net/http"
	"os"
)

var container = BuildContainer()

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	router := mux.NewRouter()

	router.HandleFunc(config.TourneyCreateUri, createTourney)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8008"
	}
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}

func createTourney(w http.ResponseWriter, r *http.Request) {

	var players []*dto.Player

	err := json.NewDecoder(r.Body).Decode(&players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var createTourneyUseCase = create_tourney.NewCommand(16, players)

	err = container.Invoke(func(createTourneyHandler *create_tourney.Handler) {
		createTourneyHandler.Handle(createTourneyUseCase)
	})

	if err != nil {
		log.Println("Error container invoke:", err)
	}

	w.WriteHeader(http.StatusOK)
}

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(func() *gorm.DB {
		dbUri := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
		db, err := gorm.Open("", dbUri)
		if err != nil {
			log.Fatal(err)
		}
		return db
	})

	err := container.Provide(func() repository.TeamRepositoryInterface {
		return &infrastructureRepo.StubRepo{}
	})
	if err != nil {
		log.Fatal(err)
	}

	err = container.Provide(func(teamRepo repository.TeamRepositoryInterface) *create_tourney.Handler {
		return create_tourney.NewHandler(&teamRepo)
	})
	if err != nil {
		log.Fatal(err)
	}

	return container
}
