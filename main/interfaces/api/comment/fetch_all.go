package comment

import (
	"golang.org/x/net/websocket"
	"log"
)

func FetchAll(ws *websocket.Conn) {
	//entity := comment.Comment{}
	type T struct {
		Msg   string
		Count float64
	}
	var data T

	websocket.JSON.Receive(ws, &data)

	log.Printf("data=%#v\n", data)

	websocket.JSON.Send(ws, data)
}
