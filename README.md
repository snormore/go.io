Go.iO - a Go push server
=====

A Go application providing a push server using sockjs and rabbitmq.

![Go.iO Communication Sequence Diagram](https://raw.github.com/snormore/go.io/master/resource/communication-sequence.png)
![Go.iO Node Sequence Diagram](https://raw.github.com/snormore/go.io/master/resource/gonode-sequence.png)

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
