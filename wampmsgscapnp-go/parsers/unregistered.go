package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Unregistered struct {
	gen *gen.Unregistered
}

func NewUnregisteredFields(g *gen.Unregistered) messages.UnregisteredFields {
	return &Unregistered{gen: g}
}

func (u *Unregistered) RequestID() int64 {
	return u.gen.RequestID()
}

func UnregisteredToCapnproto(m *messages.Unregistered) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	unregistered, err := gen.NewUnregistered(seg)
	if err != nil {
		return nil, err
	}

	unregistered.SetRequestID(m.RequestID())

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeUnregistered)}, data.Bytes()...), nil
}

func CapnprotoToUnregistered(data []byte) (*messages.Unregistered, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	ur, err := gen.ReadRootUnregistered(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewUnregisteredWithFields(NewUnregisteredFields(&ur)), nil
}
