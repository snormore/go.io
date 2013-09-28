package consumer_transport

import (
	"github.com/snormore/go.io/dispatcher/message"
	"log"
)

type RedisConsumerTransport struct {
	connection interface{}
	queue_key  string
	err        error
}

func NewRedisConsumerTransport() RedisConsumerTransport {
	log.Print("RedisConsumerTransport: initializing...")

	return RedisConsumerTransport{
		queue_key: "gopush",
	}
}

func (self RedisConsumerTransport) GetError() error {
	return self.err
}

func (self RedisConsumerTransport) Destroy() {

}

func (self RedisConsumerTransport) Listen(messages chan dispatcher_message.Message) {

}
