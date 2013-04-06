package main

import (
    "github.com/igm/sockjs-go/sockjs"
    "net/http"
    "log"
    "go.io/queue"
    "time"
    "flag"
)

func main() {
    log.Println("server started")
    sockjs.Install("/echo", SockJSHandler, sockjs.DefaultConfig)
    http.Handle("/", http.FileServer(http.Dir("./www")))
    err := http.ListenAndServe(":8080", nil)
    log.Fatal(err)
}

func SockJSHandler(session sockjs.Conn) {
    log.Println("Session created")
    consumer, err := queue.GetConsumer()
    for {
        // val, err := session.ReadMessage()
        lifetime := flag.Duration("lifetime", 0*time.Second, "lifetime of process before shutdown (0s=infinite)")
        if *lifetime > 0 {
            log.Printf("running for %s", *lifetime)
            time.Sleep(*lifetime)
        } else {
            log.Printf("running forever")
            select {}
        }
        if err != nil {
            break
        }
        go func() { session.WriteMessage([]byte("Hello!")) }()
    }
    log.Printf("shutting down")
    if err := consumer.Shutdown(); err != nil {
        log.Fatalf("error during shutdown: %s", err)
    }
    log.Println("session closed")
}
