package server

import (
	"flag"
	"os"
	"time"

	"github.com/Avyukth/streaming.service/internal/handlers"
	w "github.com/Avyukth/streaming.service/pkg/rtc"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

var (
	address     = flag.String("address", ":"+os.Getenv("PORT"), "address to listen on")
	certificate = flag.String("cert", "", "certificate file")
	key         = flag.String("key", "", "key file")
)

func Run() error {
	flag.Parse()
	if *address == ":" {
		*address = ":8080"
	}
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket", websocket.New(handlers.RoomWebsocket, websocket.Config{
		HandshakeTimeout: time.Second * 10}))
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
	app.Get("/stream/:ssuid/websocket", websocket.New(handlers.StreamWebsocket, websocket.Config{
		HandshakeTimeout: time.Second * 10}))
	app.Get("/stream/:ssuid/chat/websocket", websocket.New(handlers.StreamChatWebsocket))
	app.Get("/stream/:ssuid/viewer/websocket", websocket.New(handlers.StreamViewerWebsocket))
	app.Static("/", "./assets")

	w.Rooms = make(map[string]*w.Room)
	w.Streams = make(map[string]*w.Room)
	if *certificate != "" && *key != "" {
		app.ListenTLS(*address, *certificate, *key)
	} else {
		app.Listen(*address)
	}
	go dispatchKeyFrames()
	return app.Listen(*address)
}
func dispatchKeyFrames() {
	for range time.NewTicker(time.Second * 1).C {
		for _, room := range w.Rooms {
			room.DispatchKeyFrame()
		}
	}
}
