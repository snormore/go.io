package auth_transport

import (
	"go.io/dispatcher/client"
)

type ApiAuthTransport struct {
}

func NewApiAuthTransport() ApiAuthTransport {
	self := ApiAuthTransport{}
	return self
}

func (self *ApiAuthTransport) Authenticate(client *dispatcher_client.Client) error {
	// TODO: authenticate and stuff...
	return nil
}

func (self *ApiAuthTransport) Destroy() error {
	// TODO: authenticate and stuff...
	return nil
}
