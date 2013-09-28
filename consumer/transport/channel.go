package consumer_transport

import (
	"github.com/snormore/go.io/dispatcher/message"
	"launchpad.net/tomb"
	"log"
)

type ChannelConsumerTransport struct {
	messages chan dispatcher_message.Message
	err      error
}

func NewChannelConsumerTransport() ChannelConsumerTransport {
	log.Print("ChannelConsumerTransport: initializing...")

	return ChannelConsumerTransport{
		messages: make(chan dispatcher_message.Message),
	}
}

func (self ChannelConsumerTransport) GetError() error {
	return self.err
}

func (self ChannelConsumerTransport) Destroy() {
	close(self.messages)
}

func (self ChannelConsumerTransport) Listen(messages chan dispatcher_message.Message, t *tomb.Tomb) {
	defer t.Done()

	for {
		select {
		case msg := <-self.messages:
			messages <- msg
		case <-t.Dying():
			return
		}
	}
}
