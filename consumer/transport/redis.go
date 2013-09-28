package consumer_transport

import (
	"encoding/json"
	"github.com/snormore/go.io/dispatcher/message"
	"log"
)

var redisMessages = make(chan dispatcher_message.Message)

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
	for msg := range redisMessages {
		messages <- msg
	}
}

func (self RedisConsumerTransport) CreateDispatcherMessage(encodedMsg []byte) (dispatcher_message.Message, error) {
	var msg dispatcher_message.Message
	err := json.Unmarshal(encodedMsg, &msg)
	if err != nil {
		// TODO: log and stuff...
		return dispatcher_message.NewErrorMessage(err), err
	}
	return msg, nil
}
