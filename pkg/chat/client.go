package chat


import (
	"time"
	"github.com/fasthttp/websocket"
)


const(
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
)
type Client struct {
	Hub *Hub
	Conn  *websocket.Conn
	Send chan []byte
}


func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	}
	c.Conn.ReadMessage()
}

func(c *Client)WritePump(){

}

func PeerChatConnection(){
	
}
