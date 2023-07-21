package webrtc

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

func RoomConnection(c *websocket.Conn, p *Peers) (conn *webrtc.PeerConnection, err error) {
	var config webrtc.Configuration

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Println(err)
		return
	}

	newPeer := PeerConnectionState{
		PeerConnection: peerConnection,
		websocket:      &ThreadSafeWriter{},
		Conn:           c,
		Mutex:          sync.Mutex{},
	}
}
