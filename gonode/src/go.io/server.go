package main

import (
	"go.io/consumer"
	"go.io/dispatcher"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Go.iO node starting...")

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT)

	// messageChannel := make(chan string)
	// clientChannel := make(chan client.Client)

	_dispatcher := dispatcher.NewDispatcher()
	go _dispatcher.Listen()

	_consumer := consumer.NewConsumer()
	go _consumer.Listen()

	<-signalChannel
	_dispatcher.Destroy()
	_consumer.Destroy()
	// if env.Config.Verbosity > 0 {
	//     log.Println("Reports Firehose Consumer - node ending")
	// }
	log.Println("Go.iO node finished.")
}
