package message

import (
	"time"

	"github.com/pborman/uuid"

	"github.com/mccuskero/go-restful-sandbox/pkg/randomgen"
)

type Message struct {
	Id      string
	UserID	  string
	Payload   string
	Timestamp int64
}

func CreateMessge() *Message {
	event := &Message{
		Id:      uuid.NewRandom().String(),
		UserID:      uuid.NewRandom().String(),
		Payload:   randomgen.AlphaNumericString(20),
		Timestamp: time.Now().Unix(),
	}

	return event
}

func CreateMessages(numMsgs int) []*Message {

	messages := make([]*Message, numMsgs)

	for i := range messages {
		messages[i] = CreateMessge()
	}

	return messages
}
