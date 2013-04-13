package main

import (
	"go.io/consumer"
	"go.io/dispatcher"
	"go.io/dispatcher/message"
	"go.io/env"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("Go.iO node starting...")
	env.NodeStartedAt = time.Now()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT)

	messageChannel := make(chan dispatcher_message.Message)

	_dispatcher := dispatcher.NewDispatcher()
	go _dispatcher.Listen(messageChannel)

	_consumer := consumer.NewConsumer()
	go _consumer.Listen(messageChannel)

	<-signalChannel
	flushMessagesChannel(messageChannel)
	_dispatcher.Destroy()
	_consumer.Destroy()
	// if env.Config.Verbosity > 0 {
	//     log.Println("Reports Firehose Consumer - node ending")
	// }
	log.Println("Go.iO node finished.")
}

func flushMessagesChannel(ch chan dispatcher_message.Message) {
	select {
	case <-ch:
	default:
		break
	}
}
