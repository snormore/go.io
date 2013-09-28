package auth

import (
	"github.com/snormore/go.io/auth/transport"
)

func NewAuthTransport() auth_transport.AuthTransport {
	t := auth_transport.BasicAuthTransport{}
	return auth_transport.AuthTransport(&t)
}
