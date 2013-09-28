package auth_transport

import (
	"github.com/snormore/go.io/dispatcher/client"
)

type BasicAuthTransport struct {
}

func NewBasicAuthTransport() BasicAuthTransport {
	self := BasicAuthTransport{}
	return self
}

func (self *BasicAuthTransport) Authenticate(client *dispatcher_client.Client) error {
	// TODO: authenticate and stuff...
	return nil
}

func (self *BasicAuthTransport) Destroy() error {
	// TODO: authenticate and stuff...
	return nil
}
