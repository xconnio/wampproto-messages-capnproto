package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Registered struct {
	gen *gen.Registered
}

func NewRegisteredFields(g *gen.Registered) messages.RegisteredFields {
	return &Registered{gen: g}
}

func (r *Registered) RequestID() int64 {
	return r.gen.RequestID()
}

func (r *Registered) RegistrationID() int64 {
	return r.gen.RegistrationID()
}

func RegisteredToCapnproto(m *messages.Registered) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	registered, err := gen.NewRegistered(seg)
	if err != nil {
		return nil, err
	}

	registered.SetRequestID(m.RequestID())
	registered.SetRegistrationID(m.RegistrationID())

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeRegistered)}, data.Bytes()...), nil
}

func CapnprotoToRegistered(data []byte) (*messages.Registered, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	registered, err := gen.ReadRootRegistered(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewRegisteredWithFields(NewRegisteredFields(&registered)), nil
}
