package client

import (
	"bufio"
	"log"
	"net"

	"github.com/VanBur/tcp-chat/internal/message"
)

func New(connection net.Conn) *Client {
	client := &Client{
		Register:  make(chan string),
		Leave:     make(chan string),
		Broadcast: make(chan *message.Message),
		Outgoing:  make(chan *message.Message),

		reader:    bufio.NewReader(connection),
		writer:    bufio.NewWriter(connection),
		isOffline: false,
	}

	client.Listen()

	return client
}

type Client struct {
	Register  chan string
	Leave     chan string
	Broadcast chan *message.Message
	Outgoing  chan *message.Message

	reader    *bufio.Reader
	writer    *bufio.Writer
	isOffline bool
}

func (c *Client) Read() {
	for {
		if c.isOffline {
			break
		}

		msg, err := message.ParseMessageFromReader(c.reader)
		if err != nil {
			log.Println("read message", err)
			continue
		}

		switch msg.CommandType {
		case message.Connect:
			c.Register <- msg.User
		case message.Disconnect:
			c.Leave <- msg.User
		case message.Broadcast:
			c.Broadcast <- msg
		default:
			log.Println("unknown command")
			continue
		}
	}
}

func (c *Client) Write() {
	for data := range c.Outgoing {
		msgBytes, err := data.ToBytes()
		if err != nil {
			log.Println("msg to bytes error", err)
			continue
		}

		_, err = c.writer.Write(msgBytes)
		if err != nil {
			log.Println("writer write msg error", err)
			continue
		}

		err = c.writer.Flush()
		if err != nil {
			log.Println("writer flush error", err)
		}
	}
}

func (c *Client) Listen() {
	go c.Read()
	go c.Write()
}

func (c *Client) Down() {
	c.isOffline = true
}
