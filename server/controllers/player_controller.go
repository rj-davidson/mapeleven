package controllers

import (
	"mapeleven/models"
	"net/http"
)

type PlayerInfo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Age         int    `json:"age"`
	Birth       Birth  `json:"birth"`
	Nationality string `json:"nationality"`
	Height      string `json:"height"`
	Weight      string `json:"weight"`
	Injured     bool   `json:"injured"`
	Photo       string `json:"photo"`
}

type Birth struct {
	Date    string `json:"date"`
	Place   string `json:"place"`
	Country string `json:"country"`
}

type Player struct {
	Player PlayerInfo `json:"player"`
}

type PlayerResponse struct {
	Response []Player `json:"response"`
}

type PlayerController struct {
	client      *http.Client
	playerModel *models.PlayerModel
}
