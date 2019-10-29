package main

import (
	"domain/team/repository"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/dig"
	infrastructureRepo "infrastructure/repository"
	"net/http"
	"os"
	"toury_bakcend/src/application/usecase/player/dto"
	"toury_bakcend/src/application/usecase/tourney/create_tourney"
)

var container = BuildContainer()

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/tourney/create/", createTourney)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}

func createTourney(w http.ResponseWriter, r *http.Request) {

	p := &dto.Player{"P1", 4, []uint8{1, 2}}
	var p2 = &dto.Player{"P2", 12, []uint8{3, 8}}

	var createTourneyUseCase = create_tourney.Command{16, []*dto.Player{p, p2}}

	container.Invoke(func(createTourneyHandler *create_tourney.Handler) {
		createTourneyHandler.Handle(createTourneyUseCase)
	})

	//w.WriteHeader(http.StatusOK)
}

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(func() repository.TeamRepository {
		return &infrastructureRepo.StubRepo{}
	})
	container.Provide(func(teamRepo *repository.TeamRepository) *create_tourney.Handler {
		return create_tourney.NewHandler(teamRepo)
	})

	return container
}
