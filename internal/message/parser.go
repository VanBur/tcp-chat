package message

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

func ParseMessageFromReader(r io.Reader) (*Message, error) {
	size := make([]byte, MsgSize)
	_, err := r.Read(size)
	if err != nil {
		return nil, err
	}

	payloadSize := binary.BigEndian.Uint16(size)
	payload := make([]byte, payloadSize)

	_, err = r.Read(payload)
	if err != nil {
		return nil, err
	}

	msg := &Message{}

	err = json.Unmarshal(payload, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
