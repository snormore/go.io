Go.iO - a Go push server
=====

A Go application providing a push server using a pluggable client transports (such as sockjs, socketio, SSE) and a pluggable consumer transport (such as rabbitmq, sqs, redis) [under development]

![Go.iO Communication Sequence Diagram](https://raw.github.com/snormore/go.io/master/docs/communication-sequence.png)
![Go.iO Node Sequence Diagram](https://raw.github.com/snormore/go.io/master/docs/gonode-sequence.png)

Bootstrap
=========

Download and install the latest Vagrant installer from http://downloads.vagrantup.com/
```
wget wget http://files.vagrantup.com/packages/64e360814c3ad960d810456add977fd4c7d47ce6/Vagrant.dmg
```
Now clone the repo and bootstrap your workspace
```
git clone git@github.com:snormore/go.io.git
cd go.io
vagrant up
vagrant ssh
cd /u/apps/go.io
source sbin/init-env.sh
sbin/go-get.sh
```

If you use SublimeText as your editor, check out the plugin https://github.com/DisposaBoy/GoSublime

Building
========
```
sh sbin/build.sh
```

Running
=======
```
sh sbin/run.sh
```

Building and running
====================
```
sh sbin/build-and-run.sh
```

TODO
====
- Consider an authentication strategy for Javascript clients connecting to GoNode, perhaps proxied from the master WebNode.
