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
}

// NewNetwork creates an empty network.
func NewNetwork() *Network {
	return &Network{
		vms:     make(map[int]*intcode.VM),
		msgs:    make(map[int]chan int),
		packets: make(chan packet, 128),
	}
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

// Listen runs the network, starting each VM and sending inputs and outputs between
// them as requested. Returns the X and Y coordinates sent to the nonexistent
// machine at address 255.
func (net *Network) Listen() (int, int) {
	for addr, vm := range net.vms {
		go net.listenAndRun(addr, vm)
	}

	for packet := range net.packets {
		if packet.addr == 255 {
			return packet.x, packet.y
		}

		if ch, ok := net.msgs[packet.addr]; ok {
			ch <- packet.x
			ch <- packet.y
		} else {
			log.Printf("got packet to an unregistered VM: %#v", packet)
		}
	}

	panic("never got a packet for address 255")
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
