package player_models

//
//import (
//	"context"
//	"mapeleven/db/ent"
//	"mapeleven/db/ent/club"
//	"mapeleven/db/ent/league"
//	"mapeleven/db/ent/player"
//	"mapeleven/db/ent/playerseason"
//	"mapeleven/db/ent/season"
//	_ "mapeleven/db/ent/team"
//)
//
//// CreatePlayerSeasonInput holds the required fields to create a new player season.
//type CreatePlayerSeasonInput struct {
//	Year        *string
//	TeamId      int
//	PlayerId    int
//	Appearances string
//	Goals       int
//	Assists     int
//	Saves       int
//	Position    string
//	Nationality string
//	Name        string
//	Season      int
//}
//
//type UpdatePlayerSeasonInput struct {
//	Year        *string
//	TeamId      int
//	PlayerId    int
//	Appearances string
//	Goals       int
//	Assists     int
//	Saves       int
//	Position    string
//	Nationality string
//	Name        string
//	Season      int
//}
//
//// DeletePlayerInput holds the required fields to delete a player.
//type DeletePlayerInput struct {
//	ID int
//}
//
//type PlayerSeasonModel struct {
//	client *ent.Client
//}
//
//// NewPlayerSeasonModel creates a new PlayerSeasonModel.
//func NewPlayerSeasonModel(client *ent.Client) *PlayerSeasonModel {
//	return &PlayerSeasonModel{client: client}
//}
//
//// CreatePlayerSeason creates a new player season.
//func (m *PlayerSeasonModel) CreatePlayerSeason(ctx context.Context, year, apiFootballLeagueId, teamId, playerId int) (*ent.PlayerSeason, error) {
//	// Find the season by its year
//	s, err := m.client.Season.
//		Query().
//		Where(
//			season.YearEQ(year),
//			season.HasLeagueWith(
//				league.FootballApiIdEQ(apiFootballLeagueId),
//			),
//		).
//		WithLeague().
//		Only(context.Background())
//
//	if err != nil {
//		return nil, err
//	}
//
//	// Find the club by its apiFootballId
//	c, err := m.client.Club.
//		Query().
//		Where(club.ApiFootballIdEQ(teamId)).
//		Only(context.Background())
//
//	if err != nil {
//		return nil, err
//	}
//
//	// Find the player by its ID
//	p, err := m.client.Player.
//		Query().
//		Where(player.IDEQ(playerId)).
//		Only(context.Background())
//
//	if err != nil {
//		return nil, err
//	}
//
//	// Create the PlayerSeason
//	ps, err := m.client.PlayerSeason.
//		Create().
//		SetSeason(s).
//		SetClub(c).
//		SetPlayer(p).
//		Save(ctx)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return ps, nil
//}
//
//// GetPlayerSeason retrieves a player season by its player ID and year.
//func (m *PlayerSeasonModel) GetPlayerSeason(ctx context.Context, playerId, year int) (*ent.PlayerSeason, error) {
//	ps, err := m.client.PlayerSeason.
//		Query().
//		Where(
//			playerseason.HasPlayerWith(
//				player.IDEQ(playerId),
//			),
//			playerseason.HasSeasonWith(
//				season.YearEQ(year),
//			),
//		).
//		WithPlayer().
//		WithSeason().
//		Only(ctx)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return ps, nil
//}
//
//// CreateOrUpdatePlayerSeason creates or updates a player season.
//func (m *PlayerSeasonModel) CreateOrUpdatePlayerSeason(ctx context.Context, year, playerId, appearances, goals, assists, saves int) (*ent.PlayerSeason, error) {
//	// Check if a player season record for this player and year already exists
//	existing, err := m.GetPlayerSeason(ctx, playerId, year)
//	if err != nil && !ent.IsNotFound(err) {
//		return nil, err
//	}
//
//	// If it exists, update it
//	if existing != nil {
//		return m.client.PlayerSeason.
//			UpdateOneID(existing.ID).
//			SetAppearances(appearances).
//			SetGoals(goals).
//			SetAssists(assists).
//			SetSaves(saves).
//			Save(ctx)
//	}
//
//	// Otherwise, create a new record
//	return m.client.PlayerSeason.
//		Create().
//		SetYear(year).
//		SetPlayerID(playerId).
//		SetAppearances(appearances).
//		SetGoals(goals).
//		SetAssists(assists).
//		SetSaves(saves).
//		Save(ctx)
//}
//
//// ListAll retrieves a list of all PlayerSeasons.
//func (m *PlayerSeasonModel) ListAll(ctx context.Context) ([]*ent.PlayerSeason, error) {
//	return m.client.PlayerSeason.
//		Query().
//		WithSeason(
//			func(s *ent.SeasonQuery) {
//				s.WithLeague()
//			}).WithClub().
//		All(ctx)
//}
