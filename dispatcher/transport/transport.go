package dispatcher_transport

import (
	"github.com/snormore/go.io/auth/transport"
	"github.com/snormore/go.io/dispatcher/client"
	"github.com/snormore/go.io/dispatcher/message"
	"launchpad.net/tomb"
)

type DispatcherTransport interface {
	Listen(messages chan dispatcher_message.Message, clients *dispatcher_client.Clients, t *tomb.Tomb)
	Destroy()
}

func NewDispatcherTransport(auth *auth_transport.AuthTransport) DispatcherTransport {
	// TODO: which transport to use should come from env config
	t := NewChannelDispatcherTransport(auth)
	return DispatcherTransport(t)
}
