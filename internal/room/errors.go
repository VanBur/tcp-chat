package room

import "errors"

var (
	ErrMessageIsNil       = errors.New("message is nil")
	ErrChatMessageIsEmpty = errors.New("chat message is empty")
)
