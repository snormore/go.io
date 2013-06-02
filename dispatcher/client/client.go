package dispatcher_client

type Client interface {
	SendMessage(msg string) error
}

type Clients struct {
	clients []Client
}

func NewClients(clients []Client) Clients {
	self := Clients{clients}
	return self
}

func (self *Clients) Iter() []Client {
	return self.clients
}

func (self *Clients) Add(client Client) {
	self.clients = append(self.clients, client)
}

func (self *Clients) Remove(client Client) {
	// TODO: implement this client removal method...
}
