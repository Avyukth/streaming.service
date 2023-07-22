package webrtc

import (
	"sync"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}
type Peers struct {
	ListLock    sync.RwMutex
	Connection  []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

type PeerConnectionState struct {
	PeerConnection   *webrtc.PeerConnection
	websocket        websocket
	Mutex            sync.RWMutex
	Conn             connection
	ThreadSafeWriter *ThreadSafeWriter
}

func NewPeerConnectionState() *PeerConnectionState {
	return &PeerConnectionState{
		PeerConnection: nil,
	}

}
