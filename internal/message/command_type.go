package message

type CommandType int

const (
	Connect CommandType = iota + 1
	Broadcast
	Disconnect
)

func (c CommandType) invalid() bool {
	return c < Connect || c > Disconnect
}
