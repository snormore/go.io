gopush - a Go push server
======

A Go application providing a push server using a pluggable client transports (such as sockjs, socketio, SSE) and a pluggable consumer transport (such as rabbitmq, sqs, redis) [*under development*]

![gopush sequence diagram](https://raw.github.com/snormore/gopush/master/docs/gopush-sequence.png)

If you don't have a `GOPATH` environment variable set already, then do something like:
```sh
export GOPATH=~/Go
```

Bootstrap
=========
```sh
go get github.com/snormore/gopush
cd $GOPATH/src/github.com/snormore/gopush
go get -u
```

Testing
=======
```sh
go test ./...
```

Buiding
=======
```sh
go build -a -o bin/gopush
```

TODO
====
 [ ] implement pluggable auth transport for use by the dispatcher to authentication clients
 [ ] implement router logic in dispatcher for routing messages based on specified channel/topic/key
