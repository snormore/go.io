package dispatcher_transport

import (
	"github.com/igm/sockjs-go/sockjs"
	"go.io/dispatcher/client"
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

func (self *SockjsDispatcherTransport) Listen(messageChannel chan dispatcher_message.Message, clients *dispatcher_client.Clients) {
	sockjs.Install("/sockjs", func(session sockjs.Conn) {
		self.ConnectionHandler(session, clients)
	}, sockjs.DefaultConfig)
	http.Handle("/", http.FileServer(http.Dir("./www")))
	http_err := http.ListenAndServe(":8080", nil)
	log.Fatal(http_err)
}

func (self *SockjsDispatcherTransport) Destroy() {

}

type SockjsClient struct {
	session *sockjs.Conn
}

func (self *SockjsDispatcherTransport) ConnectionHandler(session sockjs.Conn, clients *dispatcher_client.Clients) {
	client := SockjsClient{&session}
	clients.Add(client)
	log.Println("Client session created: transport=sockjs)")
	for {
		_, err := session.ReadMessage()
		if err != nil {
			break
		}
		go func() { session.WriteMessage([]byte("{\"Id\":\"PING\", \"Body\":\"Ping\", \"Error\":null}")) }()
		time.Sleep(1 * time.Second)
	}
	log.Println("Client session closed: transport=sockjs")
}

func (self SockjsClient) SendMessage(msg string) error {
	log.Printf("Sending message to client %s", self.session)
	s := self.session
	(*s).WriteMessage([]byte(msg))
	return nil
}
