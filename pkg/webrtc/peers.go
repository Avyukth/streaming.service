package webrtc

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}
type Peers struct {
	ListLock    sync.RWMutex
	Connection  []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

type PeerConnectionState struct {
	PeerConnection *webrtc.PeerConnection
	websocket      *ThreadSafeWriter
}

type ThreadSafeWriter struct {
	Conn  *websocket.Conn
	Mutex sync.Mutex
}

type webSocketMessage struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

func NewPeerConnectionState() *PeerConnectionState {
	return &PeerConnectionState{
		PeerConnection: nil,
	}

}

func (t *ThreadSafeWriter) WriteJSON(v interface{}) error {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	return t.Conn.WriteJSON(v)
}

func (p *Peers) AddTrack(t *webrtc.TrackRemote) *webrtc.TrackLocalStaticRTP {
	return nil

}
func (p *Peers) RemoveTrack(t *webrtc.TrackRemote) {
	return nil
}

func (p *Peers) SinglePeerConnection() {
	return nil
}
func (p *Peers) DispatchKeyFrame() {
	return nil
}
