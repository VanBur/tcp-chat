package server

import (
	"log"
	"net"

	"github.com/VanBur/tcp-chat/internal/room"
)

type Server struct {
	listener  net.Listener
	room      *room.Room
	stopServe bool
}

func New(network, address string) (*Server, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener:  listener,
		room:      room.New(),
		stopServe: false,
	}, nil
}

func (s *Server) Serve() {
	log.Println("start server")
	for {
		if s.stopServe {
			return
		}

		conn, err := s.listener.Accept()
		if err != nil {
			return
		}

		s.room.Joins <- conn
	}
}

func (s *Server) Stop() {
	log.Println("stop server")
	s.stopServe = true

	s.listener.Close()
}

func (s *Server) GetConnectedUsers() int {
	return len(s.room.Clients)
}
