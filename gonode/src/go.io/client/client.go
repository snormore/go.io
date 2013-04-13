package client

import (
	"go.io/client/transport"
	"log"
)

type Client struct {
	transport client_transport.ClientTransport
}

func NewClient() Client {
	transport := client_transport.NewClientTransport()
	self := Client{transport}
	return self
}

func (self *Client) Listen() {
	log.Print("Client: listening...")
	self.transport.Listen()
}

func (self *Client) Destroy() {
	log.Print("Client: destroying...")
	self.transport.Destroy()
}
