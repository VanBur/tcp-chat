package room

import "log"

func (r *Room) onLeave(name string) {
	_, isRegisteredUser := r.Clients[name]
	if !isRegisteredUser {
		log.Println("no user with name", name)
		return
	}

	delete(r.Clients, name)
}
