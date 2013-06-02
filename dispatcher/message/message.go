package dispatcher_message

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Id     string
	SentAt int64
	Body   string
	Error  string
}

func NewErrorMessage(err error) Message {
	msg := Message{}
	msg.Error = fmt.Sprintf("%s", err)
	return msg
}

func (self *Message) ToJson() string {
	bytes, err := json.Marshal(self)
	if err != nil {
		log.Println("Message: failed decoding to json: %s", err)
	}
	return fmt.Sprintf("%s", bytes)
}
