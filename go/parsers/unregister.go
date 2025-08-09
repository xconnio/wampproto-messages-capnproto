package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Unregister struct {
	gen *gen.Unregister
}

func NewUnregisterFields(g *gen.Unregister) messages.UnregisterFields {
	return &Unregister{gen: g}
}

func (u *Unregister) RequestID() uint64 {
	return u.gen.RequestID()
}

func (u *Unregister) RegistrationID() uint64 {
	return u.gen.RegistrationID()
}

func UnregisterToCapnproto(m *messages.Unregister) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	unregister, err := gen.NewRootUnregister(seg)
	if err != nil {
		return nil, err
	}

	unregister.SetRequestID(m.RequestID())
	unregister.SetRegistrationID(m.RegistrationID())

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeUnregister, data), nil
}

func CapnprotoToUnregister(data []byte) (*messages.Unregister, error) {
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	unregister, err := gen.ReadRootUnregister(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewUnregisterWithFields(NewUnregisterFields(&unregister)), nil
}
