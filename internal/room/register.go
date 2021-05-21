package room

import (
	"log"

	"github.com/VanBur/tcp-chat/internal/client"
)

func (r *Room) onRegister(name string, cli *client.Client) {
	_, isRegisteredUser := r.Clients[name]
	if isRegisteredUser {
		log.Printf("user [%s] already exists", name)
		return
	}

	r.Clients[name] = cli
}
