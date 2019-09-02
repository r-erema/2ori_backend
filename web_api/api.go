package main

import (
	"application/dto"
	"application/use_case"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/tourney/create/", createTourney)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}

func createTourney(w http.ResponseWriter, r *http.Request) {

	p1 := dto.Player{"Roma", 2}
	p2 := dto.Player{"Roma"}
	createTourneyUseCase := use_case.NewCreateTourney(8)

	w.WriteHeader(http.StatusOK)
}
