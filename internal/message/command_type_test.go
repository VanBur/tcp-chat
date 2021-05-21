package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsupportedCommand(t *testing.T) {
	tests := []struct {
		name    string
		cmd     CommandType
		invalid bool
	}{
		{
			name:    "valid",
			cmd:     Connect,
			invalid: false,
		},
		{
			name:    "unsupported",
			cmd:     12345,
			invalid: true,
		},
	}

	for _, currTest := range tests {
		assert.Equal(t, currTest.invalid, currTest.cmd.invalid(), currTest.name)
	}
}
