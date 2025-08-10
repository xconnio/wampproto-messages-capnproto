package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Interrupt struct {
	gen *gen.Interrupt
}

func NewInterruptFields(g *gen.Interrupt) messages.InterruptFields {
	return &Interrupt{gen: g}
}

func (i *Interrupt) RequestID() uint64 {
	return i.gen.RequestID()
}

func (i *Interrupt) Options() map[string]any {
	return map[string]any{}
}

func InterruptToCapnproto(m *messages.Interrupt) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	interrupt, err := gen.NewRootInterrupt(seg)
	if err != nil {
		return nil, err
	}

	interrupt.SetRequestID(m.RequestID())

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeInterrupt, data, nil), nil
}

func CapnprotoToInterrupt(data []byte) (*messages.Interrupt, error) {
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	interrupt, err := gen.ReadRootInterrupt(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewInterruptWithFields(NewInterruptFields(&interrupt)), nil
}
