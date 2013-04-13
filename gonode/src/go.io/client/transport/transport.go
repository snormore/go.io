package client_transport

type ClientTransport interface {
	Listen()
	Destroy()
}

func NewClientTransport() ClientTransport {
	// TODO: which transport to use should come from env config
	t := NewSockjsClientTransport()
	return ClientTransport(&t)
}
