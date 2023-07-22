package handlers

import (
	"fmt"

	w "github.com/Avyukth/streaming.service/pkg/webrtc"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	gguid "github.com/google/uuid"
)

type webSocketMessage struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

func RoomCreate(c *fiber.Ctx) error {
	roomUuid := gguid.New()
	return c.Redirect(fmt.Sprintf("/room/%s", roomUuid.String()))
}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}
	uuid, suuid, _ := CreateOrGetRoom(uuid)
	w.RoomConnection(c, room.Peers)
}

func CreateOrGetRoom(uuid string) (string, string, Room) {
	uuid := c.Params("uuid")
	if uuid == "" {
		uuid = guuid.New().String()
		c.Redirect(fmt.Sprintf("/room/%s", uuid))
	}
	return uuid
}

func RoomWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}

	_, _, room := CreateOrGetRoom(uuid)
	return
}

func RoomViewerWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}
}

func roomViewerConnection(c *websocket.Conn, p *w.Peers) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}
}
