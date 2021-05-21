package user

import (
	"net"

	"github.com/VanBur/tcp-chat/internal/message"
)

// User is a struct for tests with send/receive func
type User struct {
	conn net.Conn
}

func New(network, address string) (*User, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &User{conn: conn}, nil
}

func (u *User) Send(msg message.Message) error {
	msgBytes, err := msg.ToBytes()
	if err != nil {
		return err
	}

	if _, err := u.conn.Write(msgBytes); err != nil {
		return err
	}

	return nil
}

func (u *User) Read() (*message.Message, error) {
	msg, err := message.ParseMessageFromReader(u.conn)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
