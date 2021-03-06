package chat

import (
	"bytes"
	"time"
	"log"

	"github.com/fasthttp/websocket"
	 
)


const(
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space	= []byte{' '}

)

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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
	})
	for {
		_,message, err:=c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Hub.broadcast <- message
	}
}

func(c *Client)WritePump(){
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
		}()
	

}

func PeerChatConnection(c *websocket.Conn, hub *Hub){
	client:=&Client{Hub:hub, Conn:c, Send:make(chan []byte,256)}
	client.Hub.register <- client
	
	go client.WritePump()
	client.ReadPump()

}
