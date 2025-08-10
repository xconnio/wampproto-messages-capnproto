package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Registered struct {
	gen *gen.Registered
}

func NewRegisteredFields(g *gen.Registered) messages.RegisteredFields {
	return &Registered{gen: g}
}

func (r *Registered) RequestID() uint64 {
	return r.gen.RequestID()
}

func (r *Registered) RegistrationID() uint64 {
	return r.gen.RegistrationID()
}

func RegisteredToCapnproto(m *messages.Registered) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	registered, err := gen.NewRootRegistered(seg)
	if err != nil {
		return nil, err
	}

	registered.SetRequestID(m.RequestID())
	registered.SetRegistrationID(m.RegistrationID())

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeRegistered, data, nil), nil
}

func CapnprotoToRegistered(data []byte) (*messages.Registered, error) {
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	registered, err := gen.ReadRootRegistered(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewRegisteredWithFields(NewRegisteredFields(&registered)), nil
}
