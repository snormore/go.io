package dispatcher

import (
	"go.io/dispatcher/transport"
	"log"
)

type Dispatcher struct {
	transport dispatcher_transport.DispatcherTransport
}

func NewDispatcher() Dispatcher {
	transport := dispatcher_transport.NewDispatcherTransport()
	self := Dispatcher{transport}
	return self
}

func (self *Dispatcher) Listen() {
	log.Print("Dispatcher: listening...")
	self.transport.Listen()
}

func (self *Dispatcher) Destroy() {
	log.Print("Dispatcher: destroying...")
	self.transport.Destroy()
}
