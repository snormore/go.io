package auth_transport

import (
	"github.com/snormore/go.io/dispatcher/client"
)

type RedisSessionAuthTransport struct {
}

func NewRedisSessionAuthTransport() RedisSessionAuthTransport {
	self := RedisSessionAuthTransport{}
	return self
}

func (self *RedisSessionAuthTransport) Authenticate(client *dispatcher_client.Client) error {
	// TODO: authenticate and stuff...
	return nil
}

func (self *RedisSessionAuthTransport) Destroy() error {
	// TODO: authenticate and stuff...
	return nil
}
