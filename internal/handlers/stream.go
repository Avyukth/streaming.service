package handlers

import (
	"fmt"
	w "github.com/Avyukth/streaming.service/pkg/webrtc"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"os"
)

func Stream(c *fiber.Ctx) error {
	suuid := c.Params("suuid")
	if suuid == "" {
		c.Status(400)
		return nil
	}
	ws := "ws"

	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		ws = "wss"
	}
	w.RoomsLock.Lock()
	if _, ok := w.Rooms[suuid]; !ok {
		w.RoomsLock.Unlock()
		return c.Render("stream", fiber.Map{
			"StreamWebsocketAddr": fmt.Sprintf("%s://%s/stream/%/websocket", ws, c.Hostname(), suuid),
			"Type":                "stream",
			"ChatWebsocketAddr":   fmt.Sprintf("%s://%s/stream/%/chat/websocket", ws, c.Hostname(), suuid),
			"ViewerWebsocketAddr": fmt.Sprintf("%s://%s/stream/%/viewer/websocket", ws, c.Hostname(), suuid),
		}, "layouts/main")
	}
	w.RoomsLock.Unlock()
	return c.Render("stream", fiber.Map{"NoStream": true, "Leave": "True"}, "layouts/main")
}

func StreamWebSocket(c *websocket.Conn) {

}

func StreamViewerWebSocket(c *websocket.Conn) {

}

func viewerConnection(c *websocket.Conn) {

}
