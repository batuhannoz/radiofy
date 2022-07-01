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
