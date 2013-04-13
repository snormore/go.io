package dispatcher

import (
	"go.io/dispatcher/message"
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

func (self *Dispatcher) Listen(messageChannel chan dispatcher_message.Message) {
	log.Print("Dispatcher: listening...")
	self.transport.Listen(messageChannel)
}

func (self *Dispatcher) Destroy() {
	log.Print("Dispatcher: destroying...")
	self.transport.Destroy()
}
