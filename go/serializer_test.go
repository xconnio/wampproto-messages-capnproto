package wampprotocapnp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xconnio/wampproto-go/messages"
	wampprotocapnp "github.com/xconnio/wampproto-serializer-capnproto/go"
)

func TestHello(t *testing.T) {
	call := messages.NewCall(1, nil, "hello", nil, nil)

	serializer := wampprotocapnp.CapnprotoSerializer{}
	payload, err := serializer.Serialize(call)
	require.NoError(t, err)
	require.NotNil(t, payload)

	require.Equal(t, payload[0], uint8(messages.MessageTypeCall))

	msg, err := serializer.Deserialize(payload)
	require.NoError(t, err)
	require.NotNil(t, msg)

	recreated := msg.(*messages.Call)
	require.Equal(t, call.RequestID(), recreated.RequestID())
}
