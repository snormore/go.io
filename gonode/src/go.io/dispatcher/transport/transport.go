package dispatcher_transport

import (
	"go.io/dispatcher/client"
	"go.io/dispatcher/message"
)

type DispatcherTransport interface {
	Listen(messageChannel chan dispatcher_message.Message, clients *dispatcher_client.Clients)
	Destroy()
}

func NewDispatcherTransport() DispatcherTransport {
	// TODO: which transport to use should come from env config
	t := NewSockjsDispatcherTransport()
	return DispatcherTransport(&t)
}
