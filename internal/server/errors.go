package server

import "errors"

var (
	ErrNoAddress   = errors.New("no address")
	ErrNoResolver  = errors.New("no resolver")
	ErrBadResolver = errors.New("bad resolver")
)
