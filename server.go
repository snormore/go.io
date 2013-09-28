package main

import (
	"github.com/snormore/go.io/consumer"
	"github.com/snormore/go.io/dispatcher"
	"github.com/snormore/go.io/dispatcher/message"
	"github.com/snormore/go.io/env"
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
	flushMessageChannel(messageChannel)
	_dispatcher.Destroy()
	_consumer.Destroy()
	log.Println("Go.iO node finished.")
}

func flushMessageChannel(ch chan dispatcher_message.Message) {
	select {
	case <-ch:
	default:
		break
	}
}
