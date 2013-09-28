package dispatcher_message

import (
	"launchpad.net/gocheck"
	"testing"
)

const (
	SAMPLE_ENCODED_MESSAGE = `{"id": "12345", "sent_at": "2013-01-01 00:00:01", "body": "hello, world", "error": ""}`
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type DispatcherMessageSuite struct{}

var _ = gocheck.Suite(&DispatcherMessageSuite{})

func (s *DispatcherMessageSuite) TestParseFromEncodedJson(c *gocheck.C) {
	msg, err := Parse(SAMPLE_ENCODED_MESSAGE)
	c.Assert(err, gocheck.IsNil)
	c.Assert(msg.Id, gocheck.Equals, "12345")
	c.Assert(msg.SentAt, gocheck.Equals, "2013-01-01 00:00:01")
	c.Assert(msg.Body, gocheck.Equals, "hello, world")
	c.Assert(msg.Error, gocheck.Equals, "")
}
