package main

import (
	"application/usecase/tourney/create_tourney"
	TourneyDTO "application/usecase/tourney/dto"
	"config"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"infrastructure/di"
	"log"
	"net/http"
	"os"
)

var container = di.BuildContainer()

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	router := mux.NewRouter()

	router.HandleFunc(config.TourneyCreateUri, createTourney).Methods("POST")

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

	var tourneySettings TourneyDTO.TourneySettings

	err := json.NewDecoder(r.Body).Decode(&tourneySettings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var createTourneyUseCase = create_tourney.NewCommand(tourneySettings.TourneyTeamsCount, tourneySettings.Players)

	var tourneyDTO *TourneyDTO.TourneyDTO
	err = container.Invoke(func(createTourneyHandler *create_tourney.Handler) {
		tourneyDTO = createTourneyHandler.Handle(createTourneyUseCase)
	})
	if err != nil {
		log.Println("Error container invoke:", err)
	}

	fmt.Print()

	w.WriteHeader(http.StatusOK)
}
