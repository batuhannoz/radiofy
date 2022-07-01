package service

import (
	"backend/internal/handler/app"
	"backend/internal/store/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"strconv"
)

type ClubStore interface {
	Clubs() *[]model.Club
	CreateClub(club *model.Club) *model.Club
}

type ClubService struct {
	ClubStore
	writerByID map[uint64]*websocket.Conn
}

func NewClubService(clubStore ClubStore) *ClubService {
	writer := make(map[uint64]*websocket.Conn)
	return &ClubService{
		clubStore,
		writer,
	}
}

func (c *ClubService) Clubs() *[]app.ClubsResponse {
	clubs := c.ClubStore.Clubs()
	var clubsResponse []app.ClubsResponse
	for _, b := range *clubs {
		clubsResponse = append(clubsResponse, app.ClubsResponse{
			ID:          b.Id,
			Image:       b.Image,
			Name:        b.Name,
			Description: b.Description,
		})
	}
	return &clubsResponse
}

func (c *ClubService) CreateClub(ctx *fiber.Ctx, ClubRequest app.CreateClubRequest) *app.CreateClubResponse {
	// TODO check if user is trying to open more than one club
	user := ctx.Locals("user").(*model.User)
	CreatedClub := c.ClubStore.CreateClub(&model.Club{
		Id:          0,
		OwnerID:     user.Id,
		Name:        ClubRequest.Name,
		IsActive:    true,
		Description: ClubRequest.Description,
		Image:       ClubRequest.Image,
	})
	return &app.CreateClubResponse{
		Code:        strconv.Itoa(CreatedClub.Id),
		Name:        CreatedClub.Name,
		Image:       CreatedClub.Image,
		Description: CreatedClub.Description,
	}
}

func (c *ClubService) ListenSong(ws *websocket.Conn) error {
	user := ws.Locals("user").(*model.User)
	c.writerByID[user.Id] = ws
	return nil
}

func (c *ClubService) UptadeListenersPlayback() {

}
