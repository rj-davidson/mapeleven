package controllers

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/cmd/db/models/team_models"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"
	"context"
	"net/http"
)

type TeamIdPackage struct {
	Year                int
	ApiFootballLeagueId int
	ApiFootballClubId   int
}

type TeamController struct {
	httpClient *http.Client
	teamModel  *team_models.TeamModel
}

func NewTeamController(entClient *ent.Client) *TeamController {
	tm := team_models.NewTeamModel(entClient)
	return &TeamController{
		httpClient: http.DefaultClient,
		teamModel:  tm,
	}
}

func (tc *TeamController) EnsureTeamExists(ctx context.Context, year, apiFootballLeagueId, apiFootballClubId int) error {
	exists := tc.teamModel.Exists(ctx, year, apiFootballLeagueId, apiFootballClubId)

	if !exists {
		err := tc.fetchTeamByID(ctx, TeamIdPackage{
			Year:                year,
			ApiFootballLeagueId: apiFootballLeagueId,
			ApiFootballClubId:   apiFootballClubId,
		})
		return err
	} else {
		return nil
	}
}

func (tc *TeamController) GetTeam(ctx context.Context, year, apiFootballLeagueId, apiFootballClubId int) (*ent.Team, error) {
	return tc.teamModel.GetTeam(ctx, year, apiFootballLeagueId, apiFootballClubId)
}

func (tc *TeamController) fetchTeamByID(ctx context.Context, seasonIdentifiers TeamIdPackage) error {
	return tc.parseTeamResponse(ctx, seasonIdentifiers)
}

func (tc *TeamController) parseTeamResponse(ctx context.Context, seasonIdentifiers TeamIdPackage) error {
	return tc.upsertTeam(ctx, seasonIdentifiers)
}

func (tc *TeamController) upsertTeam(ctx context.Context, seasonIdentifiers TeamIdPackage) error {
	if tc.teamModel.Exists(ctx, seasonIdentifiers.Year, seasonIdentifiers.ApiFootballLeagueId, seasonIdentifiers.ApiFootballClubId) {
		return tc.teamModel.UpdateTeam(ctx, seasonIdentifiers.Year, seasonIdentifiers.ApiFootballLeagueId, seasonIdentifiers.ApiFootballClubId)
	} else {
		_, err := tc.teamModel.CreateTeam(ctx, seasonIdentifiers.Year, seasonIdentifiers.ApiFootballLeagueId, seasonIdentifiers.ApiFootballClubId)
		return err
	}
}
