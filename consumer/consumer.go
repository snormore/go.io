package consumer

import (
	"github.com/snormore/go.io/consumer/transport"
	"github.com/snormore/go.io/dispatcher/message"
	"log"
)

type Consumer struct {
	transport *consumer_transport.ConsumerTransport
}

func NewConsumer() Consumer {
	return Consumer{}
}

func (self *Consumer) Listen(messageChannel chan dispatcher_message.Message) {
	log.Print("Consumer: listening...")
	transport := consumer_transport.NewConsumerTransport()
	self.transport = &transport
	(*self.transport).Listen(messageChannel)
	if (*self.transport).GetError() != nil {
		log.Fatal("Consumer: error during consumer initialization: %s", (*self.transport).GetError())
	}
}

func (self *Consumer) Destroy() {
	log.Printf("Consumer: shutting down")
	(*self.transport).Destroy()
	if (*self.transport).GetError() != nil {
		log.Fatalf("Consumer: error during shutdown: %s", (*self.transport).GetError())
	}
}
