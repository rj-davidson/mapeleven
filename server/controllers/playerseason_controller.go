package controllers

//
//import (
//	"context"
//	"fmt"
//	"mapeleven/db/ent"
//	"mapeleven/models/player_models"
//	_ "net/http"
//)
//
//type PlayerSeasonIdPackage struct {
//	Year     int
//	PlayerId int
//}
//
//type PlayerSeasonController struct {
//	client            *ent.Client
//	playerSeasonModel *player_models.PlayerSeasonModel
//}
//
//func NewPlayerSeasonController(client *ent.Client) *PlayerSeasonController {
//	model := player_models.NewPlayerSeasonModel(client)
//	return &PlayerSeasonController{client: client, playerSeasonModel: model}
//}
//
//func (psc *PlayerSeasonController) EnsurePlayerSeasonExists(ctx context.Context, year, playerId int) error {
//	// Check if a player season exists for this player and year
//	fmt.Println("Initializing Player Seasons...")
//	existing, err := psc.playerSeasonModel.GetPlayerSeason(ctx, playerId, year)
//	if err != nil && !ent.IsNotFound(err) {
//		return err
//	}
//
//	// If it exists, no need to fetch and populate again
//	if existing != nil {
//		return nil
//	}
//
//	// Fetch and populate player season
//	// TODO: Fetch data from external API or other data sources
//	fmt.Println("Ensuring PlayerSeason complete.")
//	return psc.upsertPlayerSeason(ctx, PlayerSeasonIdPackage{
//		Year:     year,
//		PlayerId: playerId,
//	})
//
//}
//
//func (psc *PlayerSeasonController) upsertPlayerSeason(ctx context.Context, seasonIdentifiers PlayerSeasonIdPackage) error {
//	existing, err := psc.playerSeasonModel.GetPlayerSeason(ctx, seasonIdentifiers.PlayerId, seasonIdentifiers.Year)
//	if err != nil && !ent.IsNotFound(err) {
//		return err
//	}
//
//	// If it exists, update it
//	if existing != nil {
//		// Update existing player season
//		// TODO: Update with new data
//		_, err := psc.playerSeasonModel.CreateOrUpdatePlayerSeason(ctx, seasonIdentifiers.Year, seasonIdentifiers.PlayerId, existing.Appearances, existing.Goals, existing.Assists, existing.Saves)
//		return err
//	}
//
//	// Otherwise, create a new record
//	// TODO: Fetch and use actual data for appearances, goals, assists, and saves
//	_, err = psc.playerSeasonModel.CreateOrUpdatePlayerSeason(ctx, seasonIdentifiers.Year, seasonIdentifiers.PlayerId, 0, 0, 0, 0)
//	return err
//}
