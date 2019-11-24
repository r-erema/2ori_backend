package main

import (
	"application/usecase/player/dto"
	"application/usecase/tourney/create_tourney"
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
