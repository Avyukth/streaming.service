package server

import (
	"flag"
	"hash"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
	"golang.org/x/net/websocket"
)

var (
  address = flag.String("address,:", os.Getenv("PORT"), "address to listen on")
  certificate = flag.String("cert", "", "certificate file")
  key = flag.String("key", "", "key file")
)

func Run() {
  flag.Parse()
  if *address == ":" {
    *address = ":8080"
  }
  engine:=html.New("./views", ".html")
  app := fiber.New(fiber.Config{Views: engine})
  app.Use(cors.New())
  app.Use(logger.New())


  app.Get("/", handlers.welcome)
  app.Get("/room/create", handlers.RoomCreate)
  app.Get("/room/:uuid", handlers.Room)
  app.Get("/room/:uuid/websocket",websocket.New(handlers.RoomWebsocket, websocket.Config{
    HandshakeTimeout: time.Second * 10,}))
  app.Get("/room/:uuid/chat",handlers.RoomChat)
  app.Get("/room/:uuid/chat/websocket",websocket.RoomChatWebsocket)
  app.Get("/room/:uuid/viewer/websocket",handlers.RoomViewerWebsocket)
  app.Get("/stream/:ssuid/websocket",)
  app.Get("/stream/:ssuid/chat/websocket",websocket.StreamChat)
  app.Get("/stream/:ssuid/viewer/websocket",websocket.StreamViewer)
}

