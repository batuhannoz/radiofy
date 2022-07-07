package functions

import (
	"fmt"
	"github.com/gofiber/websocket/v2"
)

type pingPongHandler struct {
	Type string `json:"type"`
}

func PingPong(ws *websocket.Conn) {
	for {
		err := ws.WriteJSON(pingPongHandler{Type: "ping"})
		if err != nil {
			fmt.Println(err)
		}
		var pong pingPongHandler
		err = ws.ReadJSON(&pong)
		if err != nil {
			fmt.Println(err)
			ws.Close()
			return
		}
		if pong.Type != "pong" {
			ws.Close()
			return
		}
	}
}
