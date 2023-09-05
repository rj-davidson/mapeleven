package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"mapeleven/models"
	"net/http"
	"strconv"
)

type TeamSeasonsController struct {
	teamSeasonsModel *models.TeamSeasonsModel
}

func NewTeamSeasonsController(model *models.TeamSeasonsModel) *TeamSeasonsController {
	return &TeamSeasonsController{teamSeasonsModel: model}
}

func (tsc *TeamSeasonsController) CreateTeamSeason(w http.ResponseWriter, r *http.Request) {
	// Assuming you parse CreateTeamSeasonInput from request body
	var input models.CreateTeamSeasonInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	teamSeason, err := tsc.teamSeasonsModel.CreateTeamSeason(context.Background(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teamSeason)
}

func (tsc *TeamSeasonsController) UpdateTeamSeason(w http.ResponseWriter, r *http.Request) {
	// Assuming you parse UpdateTeamSeasonInput from request body
	var input models.UpdateTeamSeasonInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	teamSeason, err := tsc.teamSeasonsModel.UpdateTeamSeason(context.Background(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teamSeason)
}

func (tsc *TeamSeasonsController) DeleteTeamSeason(w http.ResponseWriter, r *http.Request) {
	// Assuming you get the teamSeasonID from request URL
	teamSeasonID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	err := tsc.teamSeasonsModel.DeleteTeamSeason(context.Background(), teamSeasonID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (tsc *TeamSeasonsController) GetTeamSeasonByID(w http.ResponseWriter, r *http.Request) {
	// Assuming you get the teamSeasonID from request URL
	teamSeasonID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	teamSeason, err := tsc.teamSeasonsModel.GetTeamSeasonByID(context.Background(), teamSeasonID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(teamSeason)
}

func (tsc *TeamSeasonsController) ListTeamSeasons(w http.ResponseWriter, r *http.Request) {
	teamSeasons, err := tsc.teamSeasonsModel.ListTeamSeasons(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(teamSeasons)
	fmt.Println("TeamSeason loaded")
}
