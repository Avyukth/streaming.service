package handlers

import (
	"fmt"
	"os"

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

	ws := "ws"

	if os.Getenv("ENVIRONMENT") == "development" {
		ws = "wss"
	}

	uuid, suuid, _ := CreateOrGetRoom(uuid)
	return c.Render("peer", fiber.Map{
		"RoomWebSocketAddr":   fmt.Sprintf("%s://%s/room%s/websocket", ws, c.Hostname(), uuid),
		"RoomLink":            fmt.Sprintf("%s://%s/room/%s", c.Protocol(), c.Hostname(), uuid),
		"ChatWebSocketAddr":   fmt.Sprintf("%s://%s/room%s/chat/websocket", ws, c.Hostname(), uuid),
		"ViewerWebSocketAddr": fmt.Sprintf("%s://%s/room%s/viewer/websocket", ws, c.Hostname(), uuid),
		"StreamLink":          fmt.Sprintf("%s://%s/stream/%s", c.Protocol(), c.Hostname(), suuid),
		"Type":                "room",
	}, "layouts/main")
}

func CreateOrGetRoom(uuid string) (string, string, *w.Room) {
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
	w.RoomConnection(c, room.Peers)
}

func RoomViewerWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}
	w.RoomsLock.Lock()
	if peer, ok := w.Rooms[uuid]; ok {
		w.RoomsLock.Unlock()
		roomViewerConnection(c, peer.Peers)
		return
	}
	w.RoomsLock.Unlock()
}

func roomViewerConnection(c *websocket.Conn, p *w.Peers) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}
}
