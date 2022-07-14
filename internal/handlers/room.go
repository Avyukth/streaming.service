package handlers

import(
	"fmt"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	guuid "github.com/google/uuid"
)


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
	uuid,suuid,_ :=CreateOrGetRoom(uuid)
}

func CreateOrGetRoom(uuid string) (string,string,Room){
	uuid := c.Params("uuid")
	if uuid == "" {
		uuid = guuid.New().String()
		c.Redirect(fmt.Sprintf("/room/%s", uuid))
	}
	return uuid
}

func RoomWebsocket(c *websocket.Conn){
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}
}
