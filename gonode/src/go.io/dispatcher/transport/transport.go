package dispatcher_transport

type DispatcherTransport interface {
	Listen()
	Destroy()
}

func NewDispatcherTransport() DispatcherTransport {
	// TODO: which transport to use should come from env config
	t := NewSockjsDispatcherTransport()
	return DispatcherTransport(&t)
}
