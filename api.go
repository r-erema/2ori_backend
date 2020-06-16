package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"toury_bakcend/src/application/usecase/team/dto"
	"toury_bakcend/src/application/usecase/team/get_teams"
	"toury_bakcend/src/application/usecase/tourney/create_tourney"
	TourneyDTO "toury_bakcend/src/application/usecase/tourney/dto"
	"toury_bakcend/src/config"
	"toury_bakcend/src/infrastructure/di"
)

var container = di.BuildContainer()

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	router := mux.NewRouter()

	router.HandleFunc(config.TourneyCreateUri, createTourney).Methods("POST", "OPTIONS")
	router.HandleFunc(config.GetTeamsUri, getTeams).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8008"
	}
	log.Println("2ori api started...")
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func getTeams(w http.ResponseWriter, r *http.Request) {
	//todo: make cors rules beautiful
	setupResponse(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	var teamsDTO *dto.TeamsDTO
	err := container.Invoke(func(createTourneyHandler *get_teams.Handler) {
		teamsDTO = createTourneyHandler.Handle()
	})
	if err != nil {
		log.Println("Error container invoke:", err)
	}

	teamsJson, err := json.Marshal(teamsDTO)
	if err != nil {
		log.Println("Error json marshall", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(teamsJson)
	if err != nil {
		log.Println("Error write response", err)
	}

}

func createTourney(w http.ResponseWriter, r *http.Request) {
	//todo: make cors rules beautiful
	setupResponse(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

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

	err = json.NewEncoder(w).Encode(tourneyDTO)
	if err != nil {
		log.Println("Error encode:", err)
	}
	w.WriteHeader(http.StatusOK)
}
