package client_transport

import (
    "github.com/igm/sockjs-go/sockjs"
    "log"
    "time"
    "net/http"
)

func Listen() {
    sockjs.Install("/sockjs", ConnectionHandler, sockjs.DefaultConfig)
    // http.Handle("/", http.FileServer(http.Dir("./www")))
    http_err := http.ListenAndServe(":8080", nil)
    log.Fatal(http_err)
}

type ClientConnection struct {
    session sockjs.Conn
}

func ConnectionHandler(session sockjs.Conn) {
    connection := ClientConnection{session}
    // TODO: add connection to the client connection channel
    log.Println("Client session created: transport=sockjs)")
    for {
        // val, err := session.ReadMessage()
        // if err != nil {
        //     break
        // }
        time.Sleep(5*time.Second)
        go func() { connection.session.WriteMessage([]byte("Hello!")) }()
    }
    log.Println("Client session closed: transport=sockjs")
}
