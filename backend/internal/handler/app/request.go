package app

type user struct {
	SpotifyID string `json:"spotify_id"`
	Token     string `json:"token"`
	Mail      string `json:"mail"`
	Country   string `json:"country"`
	Product   string `json:"product"`
}
