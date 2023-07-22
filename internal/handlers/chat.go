package handlers

import (
	"github.com/Avyukth/streaming.service/pkg/chat"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	guid "github.com/google/uuid"
)

func RoomChat(c *fiber.Ctx) error {
	return c.Render("chat", fiber.Map{}, "layouts/main")
}

func RoomChatWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")

	if uuid == "" {
		retturn

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
