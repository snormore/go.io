package consumer_transport

import (
	"github.com/snormore/go.io/dispatcher/message"
	"launchpad.net/tomb"
	"log"
)

type KafkaConsumerTransport struct {
	connection interface{}
	topic      string
	err        error
}

func NewKafkaConsumerTransport() KafkaConsumerTransport {
	log.Print("KafkaConsumerTransport: initializing...")

	return KafkaConsumerTransport{}
}

func (self KafkaConsumerTransport) GetError() error {
	return self.err
}

func (self KafkaConsumerTransport) Destroy() {

}

func (self KafkaConsumerTransport) Listen(messages chan dispatcher_message.Message, t *tomb.Tomb) {

}
