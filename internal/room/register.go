package room

import (
	"log"

	"chat/internal/client"
)

func (r *Room) onRegister(name string, cli *client.Client) {
	_, isRegisteredUser := r.Clients[name]
	if isRegisteredUser {
		log.Printf("user [%s] already exists", name)
		return
	}

	r.Clients[name] = cli
}
