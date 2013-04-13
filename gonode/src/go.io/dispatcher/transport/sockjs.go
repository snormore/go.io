package dispatcher_transport

import (
	"github.com/igm/sockjs-go/sockjs"
	"go.io/dispatcher/message"
	"log"
	"net/http"
	"time"
)

type SockjsDispatcherTransport struct {
}

func NewSockjsDispatcherTransport() SockjsDispatcherTransport {
	return SockjsDispatcherTransport{}
}

func (self *SockjsDispatcherTransport) Listen(messageChannel chan dispatcher_message.Message) {
	sockjs.Install("/sockjs", ConnectionHandler, sockjs.DefaultConfig)
	http.Handle("/", http.FileServer(http.Dir("./www")))
	http_err := http.ListenAndServe(":8080", nil)
	log.Fatal(http_err)
}

func (self *SockjsDispatcherTransport) Destroy() {

}

type ClientConnection struct {
	session *sockjs.Conn
}

func ConnectionHandler(session sockjs.Conn) {
	// connection := ClientConnection{&session}
	// client := NewClient(connection)
	// self.clientChannel <- client
	log.Println("Client session created: transport=sockjs)")
	for {
		// TODO: need goroutine health listener; basically just sessionReadMessage 
		// on all connections in channel, wait to see if connection is closed and
		// remove from channel if it is
		// val, err := session.ReadMessage()
		// if err != nil {
		//     break
		// }
		time.Sleep(5 * time.Second)
		go func() { session.WriteMessage([]byte("{\"message\":\"Hello, Go!\", \"error\":null}")) }()
	}
	log.Println("Client session closed: transport=sockjs")
}

func (self *SockjsDispatcherTransport) SendMessage(message string) {
	// client.connection.session.WriteMessage([]byte("{\"message\":\"Hello, Go!\", \"error\":null}"))
}
