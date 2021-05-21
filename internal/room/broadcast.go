package room

import (
	"chat/internal/message"
	"fmt"
	"log"
)

func onBroadcastValidation(msg *message.Message) error {
	if msg == nil {
		return ErrMessageIsNil
	}
	if msg.Msg == nil {
		return ErrChatMessageIsEmpty
	}
	return msg.Msg.Validation()
}

func (r *Room) onBroadcast(msg *message.Message) {
	if err := onBroadcastValidation(msg); err != nil {
		log.Printf("onBroadcastValidation error %q", err)
		return
	}

	_, ok := r.Clients[msg.User]
	if !ok {
		log.Printf("no user with name %s", msg.User)
		return
	}

	if msg.Msg.To != "" {
		// send to selected user
		if err := r.sendToUser(msg); err != nil {
			log.Printf("sendToUser error : %s", err)
		}

		return
	}

	// send to all users
	r.sendToAllUsers(msg)
}

func (r *Room) sendToUser(msg *message.Message) error {
	// send to selected user
	toCli, ok := r.Clients[msg.Msg.To]
	if !ok {
		return fmt.Errorf("no user with name %s", msg.Msg.To)
	}
	toCli.Outgoing <- msg

	return nil
}

func (r *Room) sendToAllUsers(msg *message.Message) {
	for cliName, currCli := range r.Clients {
		// don't send message back to author
		if cliName == msg.User {
			continue
		}

		currCli.Outgoing <- msg
	}
}
