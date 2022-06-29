package service

import (
	"backend/internal/functions"
	"backend/internal/handler/app"
	"backend/internal/store/model"
	"github.com/gofiber/fiber/v2"
)

type ClubStore interface {
	Clubs() *model.Club
	CreateClub(club *model.Club) *model.Club
}

type ClubService struct {
	ClubStore
}

func NewClubService(clubStore ClubStore) *ClubService {
	return &ClubService{
		clubStore,
	}
}

func (c *ClubService) Clubs() *app.ClubsResponse {
	return nil
}

func (c *ClubService) CreateClub(ctx *fiber.Ctx, ClubRequest app.CreateClubRequest) *app.CreateClubResponse {
	// TODO check if user is trying to open more than one club
	user := ctx.Locals("user").(*model.User)
	CreatedClub := c.ClubStore.CreateClub(&model.Club{
		Id:          0,
		OwnerID:     user.Id,
		Name:        ClubRequest.Name,
		Description: ClubRequest.Description,
		Image:       ClubRequest.Image,
		ClubCode:    functions.EncodeToString(6),
	})
	return &app.CreateClubResponse{
		Code:        CreatedClub.ClubCode,
		Name:        CreatedClub.Name,
		Image:       CreatedClub.Image,
		Description: CreatedClub.Description,
	}
}
