Go.iO - a Go push server
=====

A Go application providing a push server using sockjs and rabbitmq.

![Go.iO Communication Sequence Diagram](https://raw.github.com/snormore/go.io/master/resource/communication-sequence.png)
![Go.iO Node Sequence Diagram](https://raw.github.com/snormore/go.io/master/resource/gonode-sequence.png)

Bootstrap
=========
```
source sbin/init-env.sh
sbin/bootstrap.sh

or do it yourself:
source sbin/init-env.sh
wget vagrant.dmg
```
...

TODO
====
- Consider an authentication strategy for Javascript clients connecting to GoNode, perhaps proxied from the master WebNode.
