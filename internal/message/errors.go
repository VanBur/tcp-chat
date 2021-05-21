package message

import "errors"

var (
	ErrNoMessage              = errors.New("no message")
	ErrNoDestination          = errors.New("no destination")
	ErrNoUser                 = errors.New("no user")
	ErrUnsupportedCommandType = errors.New("unsupported command type")
	ErrOversizeData           = errors.New("oversize data")
)
