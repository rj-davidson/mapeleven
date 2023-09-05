package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"mapeleven/models"
	"net/http"
)

type StandingsController struct {
	client         *http.Client
	standingsModel *models.StandingsModel
}

func NewStandingsController(standingsModel *models.StandingsModel) *StandingsController {
	return &StandingsController{
		client:         http.DefaultClient,
		standingsModel: standingsModel,
	}
}

func (sc *StandingsController) CreateStandings(w http.ResponseWriter, r *http.Request) {
	var input models.CreateStandingsInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create the standings
	_, err = sc.standingsModel.CreateStandings(context.Background(), input)
	if err != nil {
		http.Error(w, "Could not create standings", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Standings created successfully")
}

func (sc *StandingsController) UpdateStandings(w http.ResponseWriter, r *http.Request) {
	var input models.UpdateStandingsInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	standings, err := sc.standingsModel.UpdateStandings(context.Background(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(standings)
}

func (sc *StandingsController) ListStandings(w http.ResponseWriter, r *http.Request) {
	standings, err := sc.standingsModel.ListStandings(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(standings)
	fmt.Println("Standings loaded")
}

//func (sc *StandingsController) GetStandingsDetails(w http.ResponseWriter, r *http.Request) {
//	// Assuming you have a way to get the ID from the request, for example, from the URL params
//	id := // Retrieve ID from request
//
//	standings, err := sc.standingsModel.GetStandingsByID(context.Background(), id)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	json.NewEncoder(w).Encode(standings)
//}
