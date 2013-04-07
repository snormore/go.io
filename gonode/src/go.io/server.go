package main

import (
    "log"
    "go.io/message"
    "go.io/client"
)

func main() {
    log.Println("server started")

    // initialize message listener/consumer
    message_listener := message.NewMessageListener()
    message_listener.Listen()

    // initialize client/connection listener
    client_listener := client.NewClientListener()
    client_listener.Listen()

    // TODO: we need to block here possibly, destroy should happen in MessageListener
    message_listener.Destroy()
}
