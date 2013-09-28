package dispatcher

import (
	"github.com/snormore/go.io/auth"
	"github.com/snormore/go.io/auth/transport"
	"github.com/snormore/go.io/dispatcher/client"
	"github.com/snormore/go.io/dispatcher/message"
	"github.com/snormore/go.io/dispatcher/transport"
	"launchpad.net/tomb"
	"log"
)

const (
	MAX_CLIENTS = 1014
)

type Dispatcher struct {
	transport dispatcher_transport.DispatcherTransport
	clients   *dispatcher_client.Clients
	auth      *auth_transport.AuthTransport
}

func NewDispatcher() Dispatcher {
	auth := auth.NewAuthTransport()
	transport := dispatcher_transport.NewDispatcherTransport(&auth)
	clients := dispatcher_client.NewClients(make([]dispatcher_client.Client, 0, MAX_CLIENTS))
	self := Dispatcher{transport, &clients, &auth}
	return self
}

func (self *Dispatcher) Listen(messageChannel chan dispatcher_message.Message, t *tomb.Tomb) {
	log.Print("Dispatcher: listening...")
	go self.DispatchListen(messageChannel)
	self.transport.Listen(messageChannel, self.clients, t)
}

func (self *Dispatcher) DispatchListen(messageChannel chan dispatcher_message.Message) {
	for {
		log.Printf("Dispatcher waiting to recieve messages...")
		msg := <-messageChannel
		log.Printf("Dispatcher received message: %s", msg.ToJson())
		for _, client := range self.clients.Iter() {
			log.Printf("Sending to client %s", client)
			c := dispatcher_client.Client(client)
			c.SendMessage(msg.ToJson())
		}
	}
}

func (self *Dispatcher) Destroy() {
	log.Print("Dispatcher: destroying...")
	self.transport.Destroy()
}
