package dispatcher_transport

import (
	"github.com/snormore/go.io/auth/transport"
	"github.com/snormore/go.io/dispatcher/client"
	"github.com/snormore/go.io/dispatcher/message"
	"launchpad.net/tomb"
)

type ChannelDispatcherTransport struct {
	messages chan dispatcher_message.Message
	auth     auth_transport.AuthTransport
}

func NewChannelDispatcherTransport(auth *auth_transport.AuthTransport) ChannelDispatcherTransport {
	return ChannelDispatcherTransport{
		messages: make(chan dispatcher_message.Message),
		auth:     *auth,
	}
}

func (self ChannelDispatcherTransport) Listen(messages chan dispatcher_message.Message, clients *dispatcher_client.Clients, t *tomb.Tomb) {
	defer t.Done()

	for {
		select {
		case msg := <-messages:
			self.messages <- msg
		case <-t.Dying():
			return
		}
	}
}

func (self ChannelDispatcherTransport) Destroy() {
	close(self.messages)
}
