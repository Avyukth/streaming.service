package webrtc

import ()

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}
type Peers struct {
	ListLock    sync.RwMutex
	Connection  []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

func (p *Peer) DispatchKeyFrames() {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, s := range p.streams {
		s.DispatchKeyFrames()
	}
}
