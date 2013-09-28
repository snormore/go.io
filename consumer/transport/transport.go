package consumer_transport

import (
	"github.com/snormore/go.io/dispatcher/message"
)

type ConsumerTransport interface {
	Listen(messageChannel chan dispatcher_message.Message)
	Destroy()
	CreateDispatcherMessage(encodedMsg []byte) (dispatcher_message.Message, error)
	GetError() error
}

func NewConsumerTransport() ConsumerTransport {
	transport := NewRabbitmqConsumerTransport()
	t := ConsumerTransport(&transport)
	return t
}
