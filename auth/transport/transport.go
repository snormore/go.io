package auth_transport

import (
	"github.com/snormore/go.io/dispatcher/client"
)

type AuthTransport interface {
	Authenticate(client *dispatcher_client.Client) error
	Destroy() error
}
