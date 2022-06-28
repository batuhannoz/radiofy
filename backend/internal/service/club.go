package service

import (
	"backend/internal/handler/app"
	"backend/internal/store/model"
)

type ClubStore interface {
	Clubs() *model.Club
}

type ClubService struct {
	ChatStore
}

func NewClubService(clubStore ClubStore) *ClubService {
	return &ClubService{
		clubStore,
	}
}

func (c *ClubService) Clubs() *app.ClubsResponse {
	return nil
}

func (c *ClubService) CreateClub() *app.CreateClubResponse {
	return nil
}
