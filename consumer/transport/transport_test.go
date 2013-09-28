package consumer_transport

import (
	"github.com/snormore/go.io/dispatcher/message"
	"launchpad.net/gocheck"
	"launchpad.net/tomb"
	"testing"
)

const (
	DISPATCHER_MESSAGE_CHANNEL_SIZE = 1024
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type ConsumerTransportSuite struct{}

var _ = gocheck.Suite(&ConsumerTransportSuite{})

func (s *ConsumerTransportSuite) TestListen(c *gocheck.C) {
	transport := NewChannelConsumerTransport()
	messages := make(chan dispatcher_message.Message, DISPATCHER_MESSAGE_CHANNEL_SIZE)

	var transportListenTomb tomb.Tomb
	go transport.Listen(messages, &transportListenTomb)

	transport.messages <- dispatcher_message.Message{}
	<-messages

	transportListenTomb.Kill(nil)
	transportListenTomb.Wait()
	transport.Destroy()
}
