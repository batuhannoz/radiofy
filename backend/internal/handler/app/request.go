package app

type CreateClubRequest struct {
	ID          string `json:"id"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
