package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Unregister struct {
	gen *gen.Unregister
}

func NewUnregisterFields(g *gen.Unregister) messages.UnregisterFields {
	return &Unregister{gen: g}
}

func (u *Unregister) RequestID() int64 {
	return u.gen.RequestID()
}

func (u *Unregister) RegistrationID() int64 {
	return u.gen.RegistrationID()
}

func UnregisterToCapnproto(m *messages.Unregister) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	unregister, err := gen.NewUnregister(seg)
	if err != nil {
		return nil, err
	}

	unregister.SetRequestID(m.RequestID())
	unregister.SetRegistrationID(m.RegistrationID())

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeUnregister)}, data.Bytes()...), nil
}

func CapnprotoToUnregister(data []byte) (*messages.Unregister, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	unregister, err := gen.ReadRootUnregister(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewUnregisterWithFields(NewUnregisterFields(&unregister)), nil
}
