package consumer_transport

// import (
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// 	"github.com/snormore/go.io/dispatcher/message"
// 	// "github.com/snormore/go.io/env"
// 	"github.com/streadway/amqp"
// 	"log"
// 	// "time"
// )

// TODO: these should come from an application json properties file
// var (
// 	amqp_uri      = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
// 	exchange      = flag.String("exchange", "test-exchange", "Durable, non-auto-deleted AMQP exchange name")
// 	exchange_type = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
// 	queue         = flag.String("queue", "test-queue", "Ephemeral AMQP queue name")
// 	binding_key   = flag.String("key", "test-key", "AMQP binding key")
// 	consumer_tag  = flag.String("consumer-tag", "simple-consumer", "AMQP consumer tag (should not be blank)")
// )

// type RabbitmqConsumerTransport struct {
// 	connection         *amqp.Connection
// 	connection_channel *amqp.Channel
// 	transport_id       string
// 	done               chan error
// 	err                error
// }

// func NewRabbitmqConsumerTransport() RabbitmqConsumerTransport {
// 	log.Print("RabbitmqConsumerTransport: initializing...")

// 	self := RabbitmqConsumerTransport{}
// 	self.transport_id = *consumer_tag
// 	self.done = make(chan error)

// 	log.Printf("* RabbitmqConsumerTransport: dialing %s", *amqp_uri)
// 	self.connection, self.err = amqp.Dial(*amqp_uri)
// 	if self.err != nil {
// 		self.err = fmt.Errorf("RabbitmqConsumerTransport: dial: %s", self.err)
// 		return self
// 	}

// 	log.Printf("* RabbitmqConsumerTransport: got Connection, getting Channel")
// 	self.connection_channel, self.err = self.connection.Channel()
// 	if self.err != nil {
// 		self.err = fmt.Errorf("Channel: %s", self.err)
// 		return self
// 	}

// 	log.Printf("* RabbitmqConsumerTransport: got Channel, declaring Exchange (%s)", exchange)
// 	if self.err = self.connection_channel.ExchangeDeclare(
// 		*exchange,      // name of the exchange
// 		*exchange_type, // type
// 		true,           // durable
// 		false,          // delete when complete
// 		false,          // internal
// 		false,          // noWait
// 		nil,            // arguments
// 	); self.err != nil {
// 		self.err = fmt.Errorf("* RabbitmqConsumerTransport: Exchange Declare: %s", self.err)
// 		return self
// 	}

// 	log.Printf("* RabbitmqConsumerTransport: declared Exchange, declaring Queue (%s)", queue)
// 	state, err := self.connection_channel.QueueDeclare(
// 		*queue, // name of the queue
// 		true,   // durable
// 		false,  // delete when usused
// 		false,  // exclusive
// 		false,  // noWait
// 		nil,    // arguments
// 	)
// 	self.err = err
// 	if self.err != nil {
// 		self.err = fmt.Errorf("* RabbitmqConsumerTransport: Queue Declare: %s", self.err)
// 		return self
// 	}

// 	log.Printf("* RabbitmqConsumerTransport: declared Queue (%d messages, %d consumers), binding to Exchange (key '%s')",
// 		state.Messages, state.Consumers, binding_key)

// 	if self.err != nil {
// 		log.Fatalf("%s", self.err)
// 	}
// 	return self
// }

// func (self RabbitmqConsumerTransport) GetError() error {
// 	return self.err
// }

// func (self RabbitmqConsumerTransport) Destroy() {
// 	log.Print("RabbitmqConsumerTransport: destroying...")

// 	// will close() the deliveries channel
// 	if self.err = self.connection_channel.Cancel(self.transport_id, true); self.err != nil {
// 		self.err = fmt.Errorf("* RabbitmqConsumerTransport: Consumer cancel failed: %s", self.err)
// 		return
// 	}

// 	if self.err = self.connection.Close(); self.err != nil {
// 		self.err = fmt.Errorf("* RabbitmqConsumerTransport: AMQP connection close error: %s", self.err)
// 		return
// 	}

// 	defer log.Printf("* RabbitmqConsumerTransport: AMQP shutdown OK")

// 	// wait for handle() to exit
// 	<-self.done
// }

// func (self RabbitmqConsumerTransport) Listen(messageChannel chan dispatcher_message.Message) {
// 	log.Print("RabbitmqConsumerTransport: listening...")

// 	if self.err = self.connection_channel.QueueBind(
// 		*queue,       // name of the queue
// 		*binding_key, // bindingKey
// 		*exchange,    // sourceExchange
// 		false,        // noWait
// 		nil,          // arguments
// 	); self.err != nil {
// 		self.err = fmt.Errorf("* RabbitmqConsumerTransport: Queue Bind: %s", self.err)
// 		return
// 	}

// 	log.Printf("* RabbitmqConsumerTransport: Queue bound to Exchange, starting Consume (consumer tag '%s')", self.transport_id)
// 	deliveries, err := self.connection_channel.Consume(
// 		*queue,            // name
// 		self.transport_id, // consumerTag,
// 		false,             // noAck
// 		false,             // exclusive
// 		false,             // noLocal
// 		false,             // noWait
// 		nil,               // arguments
// 	)
// 	self.err = err
// 	if self.err != nil {
// 		self.err = fmt.Errorf("* RabbitmqConsumerTransport: Queue Consume: %s", self.err)
// 		return
// 	}

// 	for d := range deliveries {
// 		msg, err := self.CreateDispatcherMessage(d.Body)
// 		if err != nil {
// 			// log and stuff...
// 		}
// 		// if time.Unix(msg.SentAt, 0).After(env.NodeStartedAt) {
// 		log.Printf("%s : %s", d.Body, msg.ToJson())
// 		messageChannel <- msg
// 		// }
// 	}
// 	log.Printf("handle: deliveries channel closed")
// 	self.done <- nil
// }

// func (self RabbitmqConsumerTransport) CreateDispatcherMessage(encodedMsg []byte) (dispatcher_message.Message, error) {
// 	var msg dispatcher_message.Message
// 	err := json.Unmarshal(encodedMsg, &msg)
// 	if err != nil {
// 		// TODO: log and stuff...
// 		return dispatcher_message.NewErrorMessage(err), err
// 	}
// 	return msg, nil
// }
