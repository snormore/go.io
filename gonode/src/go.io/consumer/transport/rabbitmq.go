package consumer_transport

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"go.io/dispatcher/message"
	"go.io/env"
	"log"
	"time"
)

// TODO: these should come from an application json properties file
var (
	amqp_uri      = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchange      = flag.String("exchange", "test-exchange", "Durable, non-auto-deleted AMQP exchange name")
	exchange_type = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	queue         = flag.String("queue", "test-queue", "Ephemeral AMQP queue name")
	binding_key   = flag.String("key", "test-key", "AMQP binding key")
	consumer_tag  = flag.String("consumer-tag", "simple-consumer", "AMQP consumer tag (should not be blank)")
)

type ConsumerTransport struct {
	connection         *amqp.Connection
	connection_channel *amqp.Channel
	transport_id       string
	done               chan error
	err                error
}

func NewConsumerTransport() ConsumerTransport {
	log.Print("ConsumerTransport: initializing...")

	self := ConsumerTransport{}
	self.transport_id = *consumer_tag
	self.done = make(chan error)

	log.Printf("* ConsumerTransport: dialing %s", *amqp_uri)
	self.connection, self.err = amqp.Dial(*amqp_uri)
	if self.err != nil {
		self.err = fmt.Errorf("ConsumerTransport: dial: %s", self.err)
		return self
	}

	log.Printf("* ConsumerTransport: got Connection, getting Channel")
	self.connection_channel, self.err = self.connection.Channel()
	if self.err != nil {
		self.err = fmt.Errorf("Channel: %s", self.err)
		return self
	}

	log.Printf("* ConsumerTransport: got Channel, declaring Exchange (%s)", exchange)
	if self.err = self.connection_channel.ExchangeDeclare(
		*exchange,      // name of the exchange
		*exchange_type, // type
		true,           // durable
		false,          // delete when complete
		false,          // internal
		false,          // noWait
		nil,            // arguments
	); self.err != nil {
		self.err = fmt.Errorf("* ConsumerTransport: Exchange Declare: %s", self.err)
		return self
	}

	log.Printf("* ConsumerTransport: declared Exchange, declaring Queue (%s)", queue)
	state, err := self.connection_channel.QueueDeclare(
		*queue, // name of the queue
		true,   // durable
		false,  // delete when usused
		false,  // exclusive
		false,  // noWait
		nil,    // arguments
	)
	self.err = err
	if self.err != nil {
		self.err = fmt.Errorf("* ConsumerTransport: Queue Declare: %s", self.err)
		return self
	}

	log.Printf("* ConsumerTransport: declared Queue (%d messages, %d consumers), binding to Exchange (key '%s')",
		state.Messages, state.Consumers, binding_key)

	if self.err != nil {
		log.Fatalf("%s", self.err)
	}
	return self
}

func (self *ConsumerTransport) GetError() error {
	return self.err
}

func (self *ConsumerTransport) Destroy() {
	log.Print("ConsumerTransport: destroying...")

	// will close() the deliveries channel
	if self.err = self.connection_channel.Cancel(self.transport_id, true); self.err != nil {
		self.err = fmt.Errorf("* ConsumerTransport: Consumer cancel failed: %s", self.err)
		return
	}

	if self.err = self.connection.Close(); self.err != nil {
		self.err = fmt.Errorf("* ConsumerTransport: AMQP connection close error: %s", self.err)
		return
	}

	defer log.Printf("* ConsumerTransport: AMQP shutdown OK")

	// wait for handle() to exit
	<-self.done
}

func (self *ConsumerTransport) Listen(messageChannel chan dispatcher_message.Message) {
	log.Print("ConsumerTransport: listening...")

	if self.err = self.connection_channel.QueueBind(
		*queue,       // name of the queue
		*binding_key, // bindingKey
		*exchange,    // sourceExchange
		false,        // noWait
		nil,          // arguments
	); self.err != nil {
		self.err = fmt.Errorf("* ConsumerTransport: Queue Bind: %s", self.err)
		return
	}

	log.Printf("* ConsumerTransport: Queue bound to Exchange, starting Consume (consumer tag '%s')", self.transport_id)
	deliveries, err := self.connection_channel.Consume(
		*queue,            // name
		self.transport_id, // consumerTag,
		false,             // noAck
		false,             // exclusive
		false,             // noLocal
		false,             // noWait
		nil,               // arguments
	)
	self.err = err
	if self.err != nil {
		self.err = fmt.Errorf("* ConsumerTransport: Queue Consume: %s", self.err)
		return
	}

	for d := range deliveries {
		msg, err := self.CreateDispatcherMessage(d.Body)
		if err != nil {
			// log and stuff...
		}
		if time.Unix(msg.SentAt, 0).After(env.NodeStartedAt) {
			log.Printf("%s : %s", d.Body, msg.ToJson())
			messageChannel <- msg
		}
	}
	log.Printf("handle: deliveries channel closed")
	self.done <- nil
}

func (self *ConsumerTransport) CreateDispatcherMessage(encodedMsg []byte) (dispatcher_message.Message, error) {
	var msg dispatcher_message.Message
	err := json.Unmarshal(encodedMsg, &msg)
	if err != nil {
		// TODO: log and stuff...
		return dispatcher_message.NewErrorMessage(err), err
	}
	return msg, nil
}
