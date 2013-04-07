package message

import (
    "go.io/message/queue"
    "log"
)

type MessageListener struct {
    consumer* message_queue.Consumer
    // TODO: we need the message channel here...
}

func NewMessageListener() MessageListener {
    return MessageListener{}
}

func (self MessageListener) Listen() {
    consumer, message_err := message_queue.GetConsumer()
    self.consumer = consumer
    // lifetime := flag.Duration("lifetime", 0*time.Second, "lifetime of process before shutdown (0s=infinite)")
    // if *lifetime > 0 {
    //     log.Printf("running for %s", *lifetime)
    //     time.Sleep(*lifetime)
    // } else {
    //     log.Printf("running forever")
    //     select {}
    // }
    if message_err != nil {
        log.Fatal("MessageListener: error during consumer initialization: %s", message_err)
    }
}

func (self MessageListener) Destroy() {
    // tear down message listener/consumer
    log.Printf("MessageListener: shutting down")
    if message_err := self.consumer.Shutdown(); message_err != nil {
        log.Fatalf("MessageListener: error during shutdown: %s", message_err)
    }
}