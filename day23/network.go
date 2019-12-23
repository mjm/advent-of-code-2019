package day23

import (
	"log"

	"github.com/mjm/advent-of-code-2019/pkg/intcode"
)

// Network is a network of interconnected Intcode VMs.
type Network struct {
	vms      map[int]*intcode.VM
	msgs     map[int]chan int
	packets  chan packet
	nextAddr int
	nat      *NAT
}

// NewNetwork creates an empty network.
func NewNetwork() *Network {
	net := &Network{
		vms:     make(map[int]*intcode.VM),
		msgs:    make(map[int]chan int),
		packets: make(chan packet, 128),
	}
	net.nat = newNAT(net)
	return net
}

// Register adds a VM to the network.
func (net *Network) Register(vm *intcode.VM) {
	address := net.nextAddr
	net.nextAddr++
	net.vms[address] = vm

	msgCh := make(chan int, 128)
	msgCh <- address
	vm.SetInputFunc(func() int {
		// can't use SetInputChan because it will block if there is nothing available
		// in the channel.
		select {
		case val := <-msgCh:
			return val
		default:
			return -1
		}
	})
	net.msgs[address] = msgCh
}

// NAT returns the NAT for this network.
func (net *Network) NAT() *NAT {
	return net.nat
}

// Listen runs the network, starting each VM and sending inputs and outputs between
// them as requested. Listen does not block: routing will continue in the background.
func (net *Network) Listen() {
	for addr, vm := range net.vms {
		go net.listenAndRun(addr, vm)
	}

	go net.routePackets()
}

func (net *Network) routePackets() {
	for packet := range net.packets {
		if ch, ok := net.msgs[packet.addr]; ok {
			ch <- packet.x
			ch <- packet.y
		} else if packet.addr != 255 {
			log.Printf("got packet to an unregistered VM: %#v", packet)
		}

		// send all packets to the NAT as well so it can monitor idleness
		net.nat.packets <- packet
	}
}

func (net *Network) listenAndRun(addr int, vm *intcode.VM) {
	go vm.MustExecute()

	for {
		addr, ok := <-vm.Output
		if !ok {
			return
		}

		x := <-vm.Output
		y := <-vm.Output

		p := packet{
			addr: addr,
			x:    x,
			y:    y,
		}

		net.packets <- p
	}
}

type packet struct {
	addr int
	x    int
	y    int
}
