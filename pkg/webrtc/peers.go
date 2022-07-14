package webrtc

func (p *Peer) DispatchKeyFrames() {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, s := range p.streams {
		s.DispatchKeyFrames()
	}
}
