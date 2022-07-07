package app

import "time"

type AuthURLResponse struct {
	Url string `json:"url"`
}

type TokenResponse struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiryTime   time.Time `json:"expiryTime"`
	RadiofyToken string    `json:"radiofyToken"`
}

type ClubsResponse struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateClubResponse struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type PlaybackResponse struct {
	AlbumID    string `json:"albumID"`
	Position   int    `json:"position"`
	SongID     string `json:"songID"`
	SongName   string `json:"songName"`
	ArtistName string `json:"artistName"`
	Image      string `json:"image"`
	ProgressMS int    `json:"progressMS"`
}

type MessageResponse struct {
	Image       string `json:"image"`
	DisplayName string `json:"displayName"`
	Message     string `json:"message"`
}

type CurrentListenersResponse struct {
	Position    int    `json:"position"`
	AlbumID     string `json:"albumID"`
	DisplayName string `json:"displayName"`
	Image       string `json:"image"`
	ArtistName  string `json:"artistName"`
	SongName    string `json:"songName"`
	SongID      string `json:"songID"`
}
