package service

import (
	"backend/internal/handler/app"
	"backend/internal/store/model"
	"fmt"
	"github.com/gofiber/websocket/v2"
)

type ChatStore interface {
	//
}

type ChatService struct {
	ChatStore
	GlobalChatClients map[uint64]*Client
}

func NewChatService(chatStore ChatStore) *ChatService {
	var Clients = make(map[uint64]*Client)
	return &ChatService{
		chatStore,
		Clients,
	}
}

func (c *ChatService) GlobalChat(ws *websocket.Conn) {
	user := ws.Locals("user").(*model.User)
	c.GlobalChatClients[user.Id] = &Client{
		ip: ws.RemoteAddr().String(),
		ws: ws,
	}
	c.GlobalChatClients[user.Id].ws.SetCloseHandler(func(code int, text string) error {
		delete(c.GlobalChatClients, user.Id)
		return nil
	})
	c.Run(c.GlobalChatClients[user.Id].ws, user)
}

func (c *ChatService) Run(ws *websocket.Conn, user *model.User) {
	for {
		var msg app.MessageRequest
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, clients := range c.GlobalChatClients {
			clients.ws.WriteJSON(app.MessageResponse{
				Image:       user.Image,
				DisplayName: user.DisplayName,
				Message:     msg.Message,
			})
		}
	}
}

func (c *ChatService) ClubChat(ws *websocket.Conn) {

}
