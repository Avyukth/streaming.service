package handlers

import (
	"fmt"
	"os"
	"time"

	w "github.com/Avyukth/streaming.service/pkg/rtc"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
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
	suuid := c.Params("suuid")
	if suuid == "" {
		return
	}
	w.RoomsLock.Lock()
	if stream, ok := w.Streams[suuid]; !ok {
		w.RoomsLock.Unlock()
		w.StreamConn(c, stream.Peers)
		return
	}
	w.RoomsLock.Unlock()
}

func StreamViewerWebSocket(c *websocket.Conn) {
	suuid := c.Params("suuid")
	if suuid == "" {
		return
	}
	w.RoomsLock.Lock()
	if stream, ok := w.Streams[suuid]; !ok {
		w.RoomsLock.Unlock()
		viewerConnection(c, stream.Peers)
		return
	}
	w.RoomsLock.Unlock()
}

func viewerConnection(c *websocket.Conn, p *w.Peers) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	defer c.Close()

	for {
		select {
		case <-ticker.C:
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write([]byte(fmt.Sprintf("%d", len(p.Connections))))
		}
	}
}
