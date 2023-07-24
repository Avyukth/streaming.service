package webrtc

import (
	"sync"

	"github.com/Avyukth/streaming.service/pkg/chat"
	"github.com/gofiber/contrib/websocket"
)

var (
	RoomsLock sync.RWMutex
	Rooms     map[string]*Room
	Streams   map[string]*Room
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}
type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionState
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
	p.ListLock.Lock()
	defer func() {
		p.ListLock.Unlock()
		p.SinglePeerConnections()
	}()
	trackLocal, err := webrtc.NewTrackLOcalStaticRTP(t.Codec().RTPCodecCapability, t.ID(), t.StreamID())
	if err != nil {
		return nil
	}
	p.TrackLocals[t.ID()] = trackLocal
	return trackLocal
}
func (p *Peers) RemoveTrack(t *webrtc.TrackRemote) {
	p.ListLock.Lock()
	defer func() {
		p.ListLock.Unlock()
		p.SinglePeerConnections()
	}()
	delete(p.TrackLocals, t.ID())
}

func (p *Peers) SinglePeerConnections() {
	// p.ListLock.Lock()
	// defer func() {
	// 	p.ListLock.Unlock()
	// 	p.DispatchKeyFrame()
	// }()
	// attemptSync:= func() (tryAgain bool) {
	// for i:= range p.Connection{

	//     if p.Connection[i].PeerConnection.ConnectionState()== webrtc.PeerConnectionStateConnected {
	//         return true
	//     }

	// }
}
func (p *Peers) DispatchKeyFrame() {
	p.ListLock.Lock()
	defer p.ListLock.Unlock()

	for i := range p.Connections {
		for _, receiver := range p.Connections[i].PeerConnection.GetReceivers() {
			if receiver.Track() == nil {
				continue
			}

			_ = p.Connections[i].PeerConnection.WriteRTCP([]rtcp.Packet{
				&rtcp.PictureLossIndication{
					MediaSSRC: uint32(receiver.Track().SSRC()),
				},
			})
		}
	}
}
