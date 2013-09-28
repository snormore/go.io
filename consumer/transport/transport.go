package consumer_transport

import (
	"github.com/snormore/go.io/dispatcher/message"
	"launchpad.net/tomb"
)

type ConsumerTransport interface {
	Listen(messages chan dispatcher_message.Message, t *tomb.Tomb)
	Destroy()
	GetError() error
}

func NewConsumerTransport() ConsumerTransport {
	transport := NewChannelConsumerTransport()
	t := ConsumerTransport(transport)
	return t
}
