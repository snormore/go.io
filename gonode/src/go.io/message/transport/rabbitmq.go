package message_transport

import (
    "flag"
    "fmt"
    "github.com/streadway/amqp"
    "log"
)

// TODO: these should come from an application json properties file
var (
    amqp_uri      = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
    exchange     = flag.String("exchange", "test-exchange", "Durable, non-auto-deleted AMQP exchange name")
    exchange_type = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
    queue        = flag.String("queue", "test-queue", "Ephemeral AMQP queue name")
    binding_key   = flag.String("key", "test-key", "AMQP binding key")
    consumer_tag  = flag.String("consumer-tag", "simple-consumer", "AMQP consumer tag (should not be blank)")
)

type MessageTransport struct {
    connection *amqp.Connection
    connection_channel *amqp.Channel
    transport_id string
    done chan error
    message_channel chan string
    err error
}

func NewMessageTransport() MessageTransport {
    log.Print("MessageTransport: initializing...")

    self := MessageTransport{}
    self.transport_id = *consumer_tag
    self.done = make(chan error)
    self.message_channel = make(chan string)

    log.Printf("* MessageTransport: dialing %s", *amqp_uri)
    self.connection, self.err = amqp.Dial(*amqp_uri)
    if self.err != nil {
        self.err = fmt.Errorf("MessageTransport: dial: %s", self.err)
        return self
    }

    log.Printf("* MessageTransport: got Connection, getting Channel")
    self.connection_channel, self.err = self.connection.Channel()
    if self.err != nil {
        self.err = fmt.Errorf("Channel: %s", self.err)
        return self
    }

    log.Printf("* MessageTransport: got Channel, declaring Exchange (%s)", exchange)
    if self.err = self.connection_channel.ExchangeDeclare(
        *exchange,     // name of the exchange
        *exchange_type, // type
        true,         // durable
        false,        // delete when complete
        false,        // internal
        false,        // noWait
        nil,          // arguments
    ); self.err != nil {
        self.err = fmt.Errorf("* MessageTransport: Exchange Declare: %s", self.err)
        return self
    }

    log.Printf("* MessageTransport: declared Exchange, declaring Queue (%s)", queue)
    state, err := self.connection_channel.QueueDeclare(
        *queue, // name of the queue
        true,  // durable
        false, // delete when usused
        false, // exclusive
        false, // noWait
        nil,   // arguments
    )
    self.err = err
    if self.err != nil {
        self.err = fmt.Errorf("* MessageTransport: Queue Declare: %s", self.err)
        return self
    }

    log.Printf("* MessageTransport: declared Queue (%d messages, %d consumers), binding to Exchange (key '%s')",
        state.Messages, state.Consumers, binding_key)

    if self.err != nil {
        log.Fatalf("%s", self.err)
    }
    return self
}

func (self *MessageTransport) GetError() error {
    return self.err
}

func (self *MessageTransport) Destroy() {
    log.Print("MessageTransport: destroying...")

    // will close() the deliveries channel
    if self.err = self.connection_channel.Cancel(self.transport_id, true); self.err != nil {
        self.err = fmt.Errorf("* MessageTransport: Consumer cancel failed: %s", self.err)
        return
    }

    if self.err = self.connection.Close(); self.err != nil {
        self.err = fmt.Errorf("* MessageTransport: AMQP connection close error: %s", self.err)
        return
    }

    defer log.Printf("* MessageTransport: AMQP shutdown OK")

    // wait for handle() to exit
    <-self.done
}

func (self *MessageTransport) Listen() {
    log.Print("MessageTransport: listening...")

    if self.err = self.connection_channel.QueueBind(
        *queue,    // name of the queue
        *binding_key,      // bindingKey
        *exchange, // sourceExchange
        false,    // noWait
        nil,      // arguments
    ); self.err != nil {
        self.err = fmt.Errorf("* MessageTransport: Queue Bind: %s", self.err)
        return
    }

    log.Printf("* MessageTransport: Queue bound to Exchange, starting Consume (consumer tag '%s')", self.transport_id)
    deliveries, err := self.connection_channel.Consume(
        *queue, // name
        self.transport_id, // consumerTag,
        false, // noAck
        false, // exclusive
        false, // noLocal
        false, // noWait
        nil,   // arguments
    )
    self.err = err
    if self.err != nil {
        self.err = fmt.Errorf("* MessageTransport: Queue Consume: %s", self.err)
        return
    }

    handle(deliveries, self.done)
}

func handle(deliveries <-chan amqp.Delivery, done chan error) {
    for d := range deliveries {
        log.Printf(
            "got %dB delivery: [%v] %s",
            len(d.Body),
            d.DeliveryTag,
            d.Body,
        )
    }
    log.Printf("handle: deliveries channel closed")
    done <- nil
}
