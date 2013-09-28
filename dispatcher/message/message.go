package dispatcher_message

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Id     string `json:"id"`
	SentAt string `json:"sent_at"`
	Body   string `json:"body"`
	Error  string `json: "error"`
}

func Parse(encodedJson string) (Message, error) {
	var msg Message
	if err := json.Unmarshal([]byte(encodedJson), &msg); err != nil {
		return NewErrorMessage(err), err
	}
	return msg, nil
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
