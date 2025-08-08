package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Interrupt struct {
	gen *gen.Interrupt
}

func NewInterruptFields(g *gen.Interrupt) messages.InterruptFields {
	return &Interrupt{gen: g}
}

func (i *Interrupt) RequestID() int64 {
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

	interrupt, err := gen.NewInterrupt(seg)
	if err != nil {
		return nil, err
	}

	interrupt.SetRequestID(m.RequestID())

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeInterrupt)}, data.Bytes()...), nil
}

func CapnprotoToInterrupt(data []byte) (*messages.Interrupt, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	interrupt, err := gen.ReadRootInterrupt(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewInterruptWithFields(NewInterruptFields(&interrupt)), nil
}
