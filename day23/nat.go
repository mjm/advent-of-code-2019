package day23

import "time"

// NAT monitors all traffic on the network.
type NAT struct {
	net           *Network
	idleDuration  time.Duration
	packets       chan packet
	last          *packet
	lastBroadcast *packet
}

func newNAT(net *Network) *NAT {
	return &NAT{
		net: net,
		// this is really slow, but shorter times seem to produce wrong results.
		idleDuration: 1000 * time.Millisecond,
		packets:      make(chan packet),
	}
}

// MonitorOnce watches all packets on the network and returns the first one sent
// to the NAT.
func (nat *NAT) MonitorOnce() (int, int) {
	for p := range nat.packets {
		if p.addr == 255 {
			return p.x, p.y
		}
	}

	panic("packets channel closed")
}

// Monitor watches all packets on the network, watching for when the network
// becomes idle. When it does, it sends the last packet sent to the NAT to
// address 0, and traffic resumes. If it would sent address 0 the same Y value
// twice in a row, it returns the packet it would have sent instead.
func (nat *NAT) Monitor(net *Network) (int, int) {
	for {
		select {
		case p := <-nat.packets:
			if p.addr == 255 {
				nat.last = &p
			}
		case <-time.After(nat.idleDuration):
			// network traffic is idle, see if we're done
			if nat.last == nil {
				continue
			}

			if nat.lastBroadcast != nil {
				if nat.lastBroadcast.y == nat.last.y {
					return nat.last.x, nat.last.y
				}
			}

			// not done yet. take the last broadcast we got and send it to
			// address 0.
			nat.lastBroadcast = nat.last
			p := *nat.last
			p.addr = 0
			net.packets <- p
		}
	}
}
