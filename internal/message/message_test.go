package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateChatMessage(t *testing.T) {
	tests := []struct {
		name   string
		msg    ChatMessage
		result []byte
		err    error
	}{
		{
			name: "valid",
			msg: ChatMessage{
				To:   "tim",
				Text: "Hello!",
			},
		},
		{
			name: "empty message",
			msg: ChatMessage{
				To:   "tim",
				Text: "",
			},
			err: ErrNoMessage,
		},
	}

	for _, currTest := range tests {
		err := currTest.msg.Validation()
		if currTest.err != nil {
			require.Error(t, err, currTest.name)
			assert.Equal(t, currTest.err, err, currTest.name)
		} else {
			require.NoError(t, err, currTest.name)
		}
	}
}
