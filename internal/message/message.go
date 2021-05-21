package message

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
)

type Message struct {
	CommandType CommandType  `json:"cmd"`
	User        string       `json:"from"`
	Msg         *ChatMessage `json:"msg,omitempty"`
}

type ChatMessage struct {
	To   string `json:"to,omitempty"`
	Text string `json:"text"`
}

func (m *ChatMessage) Validation() error {
	switch {
	case m.Text == "":
		return ErrNoMessage
	default:
		return nil
	}
}

func (m *Message) validation() error {
	if m.Msg != nil {
		if err := m.Msg.Validation(); err != nil {
			return fmt.Errorf("msg validation : %q", err)
		}
	}

	switch {
	case m.CommandType.invalid():
		return ErrUnsupportedCommandType
	case m.User == "":
		return ErrNoUser
	default:
		return nil
	}
}

func (m *Message) ToBytes() ([]byte, error) {
	if err := m.validation(); err != nil {
		return nil, fmt.Errorf("message validation %q", err)
	}

	msgPayload, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("message marshal %q", err)
	}

	// use uint16 because of MessageSize == 2
	if len(msgPayload) > math.MaxUint16 {
		return nil, ErrOversizeData
	}

	// calc msg data size
	msgPayloadSize := len(msgPayload)
	// calc data size to bytes
	msgPayloadSizeToBytes := make([]byte, MsgSize)

	// use uint16 because of MessageSize == 2
	binary.BigEndian.PutUint16(msgPayloadSizeToBytes, uint16(msgPayloadSize))

	// combine size and data to one byte slice
	result := make([]byte, 0, len(msgPayloadSizeToBytes)+msgPayloadSize)

	result = append(result, msgPayloadSizeToBytes...)
	result = append(result, msgPayload...)

	//log.Println(result)
	return result, nil
}
