package room

import (
	"chat/internal/message"
	"net"

	"chat/internal/client"
)

type Room struct {
	Clients  map[string]*client.Client
	Joins    chan net.Conn
	Shutdown chan interface{}

	register  chan string
	leave     chan string
	broadcast chan *message.Message
}

func New() *Room {
	room := &Room{
		Clients:  make(map[string]*client.Client, 0),
		Joins:    make(chan net.Conn),
		Shutdown: make(chan interface{}),

		register:  make(chan string),
		leave:     make(chan string),
		broadcast: make(chan *message.Message),
	}

	room.Listen()

	return room
}

func (r *Room) Join(conn net.Conn) {
	cli := client.New(conn)

	go func() {
		for {
			select {
			case name := <-cli.Register:
				r.onRegister(name, cli)
			case name := <-cli.Leave:
				r.onLeave(name)
			case msg := <-cli.Broadcast:
				r.onBroadcast(msg)
			}
		}
	}()
}

func (r *Room) Listen() {
	go func() {
		for {
			select {
			case conn := <-r.Joins:
				r.Join(conn)
			case <-r.Shutdown:
				return
			}
		}
	}()
}

func (r *Room) Close() {
	for name, currCli := range r.Clients {
		currCli.Down()

		delete(r.Clients, name)
	}

	r.Shutdown <- true
}
