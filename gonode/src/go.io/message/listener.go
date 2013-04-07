package message

import (
    "go.io/message/transport"
    "log"
)

type MessageListener struct {
    transport* message_transport.MessageTransport
}

func NewMessageListener() MessageListener {
    return MessageListener{}
}

func (self *MessageListener) Listen() {
    transport := message_transport.NewMessageTransport()
    self.transport = &transport
    if self.transport.GetError() != nil {
        log.Fatal("MessageListener: error during consumer initialization: %s", self.transport.GetError())
    }
}

func (self *MessageListener) Destroy() {
    log.Printf("MessageListener: shutting down")
    self.transport.Destroy()
    if self.transport.GetError() != nil {
        log.Fatalf("MessageListener: error during shutdown: %s", self.transport.GetError())
    }
}