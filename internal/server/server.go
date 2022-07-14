package server

import (
	"flag"
	"hash"
	"os"
	"time"

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
  
  app.Get("/", handlers.welcome)
  app.Get("/room/create", handlers.RoomCreate)
  app.Get("/room/:uuid", handlers.Room)
  app.Get("/room/:uuid/websocket",)
  app.Get("/room/:uuid/chat",handlers.RoomChat)
  app.Get("/room/:uuid/chat/websocket",websocket.RoomChatWebsocket)
  app.Get("/room/:uuid/viewer/websocket",handlers.RoomViewerWebsocket)
  app.Get("/stream/:ssuid/websocket",)
  app.Get("/stream/:ssuid/chat/websocket",websocket.StreamChat)
  app.Get("/stream/:ssuid/viewer/websocket",websocket.StreamViewer)
}

