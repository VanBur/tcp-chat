package server

type Conf struct {
	Resolver string
	Address  string
}

func (c *Conf) validate() error {
	switch {
	case c.Address == "":
		return ErrNoAddress
	case c.Resolver == "":
		return ErrNoResolver
	case c.Resolver == "/":
		return ErrBadResolver
	default:
		return nil
	}
}
