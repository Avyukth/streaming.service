package handlers

import (
	"github.com/Avyukth/streaming.service/pkg/chat"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	guid "github.com/google/uuid"
	"github.com/sqs/goreturns/returns"
)

func RoomChat(c *fiber.Ctx) error {
	return c.Render("chat", fiber.Map{}, "layouts/main")
}

func RoomChatWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")

	if uuid == "" {
		return

	}
	w.RoomsLock.LocK()
	room := w.Rooms[uuid]
	w.RoomsLock.Unlock()
	if room == nil {
		return
	}

	if room.Hub == nil {
		return
	}
	chat.PeerChatConn(c.Conn, room.Hub)
}

func StreamChatWebSocket(c *websocket.Conn) {
	suuid := c.Params("suuid")
	if suuid == "" {
		return
	}
	w.RoomsLock.Lock()
	if stream.Hub == nil {
		return
	}
	if stream, ok := w.Streams[suuid]; ok {

		w.RoomsLock.Unlock()
		if stream.Hub == nil {
			hub := chat.NewHub()

			stream.Hub = hub
			go hub.Run()
		}

	}
}
