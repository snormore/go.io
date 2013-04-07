package client

import (
    "go.io/client/transport"
    "log"
)

type ClientListener struct {
    connection* client_transport.ClientConnection
    // TODO: this should be the client connection channel
}

func NewClientListener() ClientListener {
    log.Print("ClientListener: initializing...")
    return ClientListener{}
}

func (self *ClientListener) Listen() {
    log.Print("ClientListener: listening...")
    client_transport.Listen()
}

func (self *ClientListener) Destroy() {
    log.Print("ClientListener: destroying...")
}