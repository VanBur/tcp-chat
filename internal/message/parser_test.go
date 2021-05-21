package message

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseMessageFromReader(t *testing.T) {
	input := []byte{
		0, 33, // size in bytes
		123, 34, 99, 109, 100, 34, 58, 49, 44, 34, // payload in bytes
		102, 114, 111, 109, 34, 58, 34, 116, 105, 109,
		34, 44, 34, 109, 115, 103, 34, 58, 110, 117,
		108, 108, 125}
	msg, err := ParseMessageFromReader(bytes.NewReader(input))
	require.NoError(t, err)
	assert.Equal(t, &Message{CommandType: Connect, User: "tim"}, msg)
}
