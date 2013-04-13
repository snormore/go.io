package main

import (
	"go.io/client"
	"go.io/consumer"
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

	client := client.NewClient()
	go client.Listen()

	consumer := consumer.NewConsumer()
	go consumer.Listen()

	// message_dispatcher := message.NewMessageDispatcher()
	// go message_dispatcher.Listen(messageChannel)

	<-signalChannel
	client.Destroy()
	consumer.Destroy()
	// if env.Config.Verbosity > 0 {
	//     log.Println("Reports Firehose Consumer - node ending")
	// }
	log.Println("Go.iO node finished.")
}
