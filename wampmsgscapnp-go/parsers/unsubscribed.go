package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Unsubscribed struct {
	gen *gen.Unsubscribed
}

func NewUnsubscribedFields(g *gen.Unsubscribed) messages.UnsubscribedFields {
	return &Unsubscribed{gen: g}
}

func (u *Unsubscribed) RequestID() int64 {
	return u.gen.RequestID()
}

func UnsubscribedToCapnproto(m *messages.Unsubscribed) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	unsubscribed, err := gen.NewUnsubscribed(seg)
	if err != nil {
		return nil, err
	}

	unsubscribed.SetRequestID(m.RequestID())

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeUnsubscribed)}, data.Bytes()...), nil
}

func CapnprotoToUnsubscribed(data []byte) (*messages.Unsubscribed, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	unsubscribed, err := gen.ReadRootUnsubscribed(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewUnsubscribedWithFields(NewUnsubscribedFields(&unsubscribed)), nil
}
