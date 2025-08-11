package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Unregistered struct {
	gen *gen.Unregistered
}

func NewUnregisteredFields(g *gen.Unregistered) messages.UnregisteredFields {
	return &Unregistered{gen: g}
}

func (u *Unregistered) RequestID() uint64 {
	return u.gen.RequestID()
}

func UnregisteredToCapnproto(m *messages.Unregistered) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	unregistered, err := gen.NewRootUnregistered(seg)
	if err != nil {
		return nil, err
	}

	unregistered.SetRequestID(m.RequestID())

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeUnregistered, data, nil), nil
}

func CapnprotoToUnregistered(data []byte) (*messages.Unregistered, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	ur, err := gen.ReadRootUnregistered(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewUnregisteredWithFields(NewUnregisteredFields(&ur)), nil
}
