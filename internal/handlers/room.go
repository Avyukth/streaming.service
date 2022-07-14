package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
	"os"
	w "streaming.service/pkg/webrtc"
	"time"
)

type webSocketMessage struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

func RoomCreate(c *fiber.Ctx) error {
	uuid := guuid.New()
	return c.Redirect(fmt.Sprintf("/room/%s", uuid.String()))
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
