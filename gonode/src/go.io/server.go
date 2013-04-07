package main

import (
    "log"
    "go.io/message"
    "go.io/client"
)

func main() {
    log.Println("server started")

    message_listener := message.NewMessageListener()
    message_listener.Listen()

    client_listener := client.NewClientListener()
    client_listener.Listen()

    // forever...
    select{}

    message_listener.Destroy()
    client_listener.Destroy()
}
