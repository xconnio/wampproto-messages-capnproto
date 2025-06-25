package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Cancel struct {
	gen *gen.Cancel
}

func NewCancelFields(g *gen.Cancel) messages.CancelFields {
	return &Cancel{gen: g}
}

func (c *Cancel) RequestID() int64 {
	return c.gen.RequestID()
}

func (c *Cancel) Options() map[string]any {
	return map[string]any{}
}

func CancelToCapnproto(m *messages.Cancel) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	c, err := gen.NewCancel(seg)
	if err != nil {
		return nil, err
	}

	c.SetRequestID(m.RequestID())

	var buf bytes.Buffer
	if err := capnp.NewEncoder(&buf).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeCancel)}, buf.Bytes()...), nil
}

func CapnprotoToCancel(data []byte) (*messages.Cancel, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	cancel, err := gen.ReadRootCancel(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewCancelWithFields(NewCancelFields(&cancel)), nil
}
