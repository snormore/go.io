package main

import (
    "log"
    "go.io/message"
    "go.io/client"
)

func main() {
    log.Println("server started")

    client_listener := client.NewClientListener()
    go client_listener.Listen();

    message_listener := message.NewMessageListener()
    go message_listener.Listen()

    // TODO: instead of waiting forever, wait for 'done' channel to return
    log.Print("Main: waiting for client and message listeners to signal exit...")
    select{}

    message_listener.Destroy()
    client_listener.Destroy()
}
