package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Cancel struct {
	gen *gen.Cancel
}

func NewCancelFields(g *gen.Cancel) messages.CancelFields {
	return &Cancel{gen: g}
}

func (c *Cancel) RequestID() uint64 {
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

	c, err := gen.NewRootCall(seg)
	if err != nil {
		return nil, err
	}

	c.SetRequestID(m.RequestID())

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeCancel, data, nil), nil
}

func CapnprotoToCancel(data []byte) (*messages.Cancel, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	cancel, err := gen.ReadRootCancel(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewCancelWithFields(NewCancelFields(&cancel)), nil
}
